package casbin

import (
	"fmt"
	"gin-casbin-admin/pkg/log"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewCasbinEnforcer(conf *viper.Viper, db *gorm.DB, logger *log.Logger) (*casbin.Enforcer, error) {
	watcher, err := rediswatcher.NewWatcher(conf.GetStringSlice("data.redis.addrs")[0], rediswatcher.WatcherOptions{
		Options:    redis.Options{Password: conf.GetString("data.redis.password")},
		Channel:    "/casbin",
		IgnoreSelf: false,
	})
	if err != nil {
		fmt.Println("create redis watcher error: ", err)
		return nil, err
	}
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, "", "casbin_rule")
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromString(conf.GetString("casbin_config.model_text"))
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	if err := e.SetWatcher(watcher); err != nil {
		return nil, err
	}
	_ = watcher.SetUpdateCallback(func(s string) {
		logger.Info("watcher update: ", zap.String("msg", s))
	})
	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
