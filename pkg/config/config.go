// package config
//
// import (
//
//	"fmt"
//	"github.com/spf13/viper"
//	"os"
//	"time"
//
// )
//
//	type Config struct {
//		Server ServerConfig `mapstructure:"server"`
//		JWT    JWTConfig    `mapstructure:"jwt"`
//		Redis  RedisConfig  `mapstructure:"redis"`
//	}
//
//	type ServerConfig struct {
//		Port int    `mapstructure:"port"`
//		Mode string `mapstructure:"mode"`
//	}
//
//	type JWTConfig struct {
//		AccessSecret  string        `mapstructure:"access_secret"`
//		RefreshSecret string        `mapstructure:"refresh_secret"`
//		AccessExpire  time.Duration `mapstructure:"access_expire"`
//		RefreshExpire time.Duration `mapstructure:"refresh_expire"`
//		Issuer        string        `mapstructure:"issuer"`
//	}
//
//	type RedisConfig struct {
//		Addr     string `mapstructure:"addr"`
//		Password string `mapstructure:"password"`
//		Database       int    `mapstructure:"db"`
//	}
//
//	type LogConfig struct {
//		LogLevel    string `mapstructure:"log_level"`
//		Encoding    string `mapstructure:"encoding"`
//		LogFileName string `mapstructure:"log_file_name"`
//		MaxSize     int    `mapstructure:"max_size"`
//		MaxAge      int    `mapstructure:"max_age"`
//		MaxBackups  int    `mapstructure:"max_backups"`
//	}
//
//	func NewConfig(p string) *viper.Viper {
//		envConf := os.Getenv("APP_CONF")
//		if envConf == "" {
//			envConf = p
//		}
//		fmt.Println("load conf file:", envConf)
//		return getConfig(envConf)
//	}
//
//	func getConfig(path string) *viper.Viper {
//		conf := viper.New()
//		conf.SetConfigFile(path)
//		err := conf.ReadInConfig()
package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Database DatabaseConfig `mapstructure:"database"`
	Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	Host string `mapstructure:"host"`
}

type JWTConfig struct {
	AccessSecret  string        `mapstructure:"access_secret"`
	RefreshSecret string        `mapstructure:"refresh_secret"`
	AccessExpire  time.Duration `mapstructure:"access_expire"`
	RefreshExpire time.Duration `mapstructure:"refresh_expire"`
	Issuer        string        `mapstructure:"issuer"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

func (c *DatabaseConfig) GetDSN() string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
		)
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			c.Host,
			c.Username,
			c.Password,
			c.Database,
			c.Port,
		)
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
		)
	}
}

type LogConfig struct {
	LogLevel    string `mapstructure:"log_level"`
	Encoding    string `mapstructure:"encoding"`
	LogFileName string `mapstructure:"log_file_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
	Compress    bool   `mapstructure:"compress"`
}

func NewConfig(p string) (*Config, error) {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)

	v := viper.New()
	v.SetConfigFile(envConf)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &cfg, nil
}
