package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/urfave/cli/v2"
	"github.com/wxlbd/gin-casbin-admin/cmd/server/wire"
	_ "github.com/wxlbd/gin-casbin-admin/docs" // 导入 swagger docs
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"go.uber.org/zap"
)

// @title Gin-Casbin-Admin API
// @version 1.0
// @description 基于 Gin + Casbin 的权限管理系统
// @termsOfService http://swagger.io/terms/

// @contact.name wxl
// @contact.url https://github.com/wxlbd
// @contact.email gopher095@gmail.com

// @license.name MIT
// @license.url https://github.com/wxlbd/gin-casbin-admin/blob/main/LICENSE

// @host localhost:8080
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

//go:embed migrations/*.sql
var migrationFS embed.FS

var (
	configFile string
	conf       *config.Config
	logger     *log.Logger
)

func main() {
	app := &cli.App{
		Name:  "gin-casbin-admin",
		Usage: "一个基于 Gin + Casbin 的权限管理系统",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "configs/config.yaml",
				Usage:       "配置文件路径",
				Destination: &configFile,
			},
		},
		Before: func(c *cli.Context) error {
			// 初始化配置和日志
			var err error
			conf, err = config.NewConfig(configFile)
			if err != nil {
				return err
			}
			logger = log.NewLog(&conf.Log)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "启动服务",
				Action:  startServer,
			},
			{
				Name:  "migrate",
				Usage: "数据库迁移",
				Subcommands: []*cli.Command{
					{
						Name:   "up",
						Usage:  "执行向上迁移",
						Action: migrateUp,
					},
					{
						Name:   "down",
						Usage:  "执行向下迁移（回滚）",
						Action: migrateDown,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error("app run error", zap.Error(err))
	}
}

func startServer(c *cli.Context) error {
	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		return err
	}
	defer cleanup()

	// 创建 http.Server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port),
		Handler: app,
	}

	// 创建一个通道来接收系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.Server.Host, conf.Server.Port)))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server run error", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	// 等待退出信号
	<-quit
	logger.Info("shutting down server...")

	// 创建一个5秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", zap.Error(err))
		return err
	}

	logger.Info("server exited")
	return nil
}

func getMigrate() (*migrate.Migrate, error) {
	// 首先确保数据库存在
	if err := ensureDBExists(); err != nil {
		return nil, fmt.Errorf("failed to ensure database exists: %w", err)
	}

	// 使用嵌入的迁移文件创建 source.Driver
	d, err := iofs.New(migrationFS, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to create iofs driver: %w", err)
	}

	// 构建数据库 DSN
	var dsn string
	switch conf.Database.Driver {
	case "mysql":
		dsn = fmt.Sprintf("mysql://%s", conf.Database.GetDSN())
	case "postgres":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Database)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", conf.Database.Driver)
	}

	// 创建 migrate 实例
	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m, nil
}

func ensureDBExists() error {
	dbConfig := conf.Database

	switch dbConfig.Driver {
	case "mysql":
		return ensureMySQLDBExists(&dbConfig)
	case "postgres":
		return ensurePostgreSQLDBExists(&dbConfig)
	default:
		return fmt.Errorf("unsupported database driver: %s", dbConfig.Driver)
	}
}

func ensureMySQLDBExists(dbConfig *config.DatabaseConfig) error {
	// 构建不带数据库名的 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port)

	// 连接 MySQL（不指定数据库）
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to mysql: %w", err)
	}
	defer db.Close()

	// 创建数据库（如果不存在）
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci", dbConfig.Database)
	if _, err := db.Exec(createDBSQL); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	logger.Info("database created or already exists", zap.String("database", dbConfig.Database))
	return nil
}

func ensurePostgreSQLDBExists(dbConfig *config.DatabaseConfig) error {
	// PostgreSQL 需要先连接到 postgres 数据库或 template1 数据库
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=postgres sslmode=disable",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer db.Close()

	// 检查数据库是否存在
	var exists bool
	checkSQL := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)"
	if err := db.QueryRow(checkSQL, dbConfig.Database).Scan(&exists); err != nil {
		return fmt.Errorf("failed to check database existence: %w", err)
	}

	if !exists {
		// 创建数据库
		createDBSQL := fmt.Sprintf("CREATE DATABASE \"%s\"", dbConfig.Database)
		if _, err := db.Exec(createDBSQL); err != nil {
			return fmt.Errorf("failed to create database: %w", err)
		}
		logger.Info("database created", zap.String("database", dbConfig.Database))
	} else {
		logger.Info("database already exists", zap.String("database", dbConfig.Database))
	}

	return nil
}

func migrateUp(c *cli.Context) error {
	m, err := getMigrate()
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate up: %w", err)
	}
	logger.Info("database migration up succeeded")
	return nil
}

func migrateDown(c *cli.Context) error {
	m, err := getMigrate()
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate down: %w", err)
	}
	logger.Info("database migration down succeeded")
	return nil
}
