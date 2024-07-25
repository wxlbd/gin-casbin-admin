package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewCasbinEnforcer(conf *viper.Viper, db *gorm.DB) (*casbin.Enforcer, error) {
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
	if err := e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
