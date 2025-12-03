//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wxlbd/gin-casbin-admin/internal/application/captcha"
	"github.com/wxlbd/gin-casbin-admin/internal/application/dict"
	"github.com/wxlbd/gin-casbin-admin/internal/application/menu"
	"github.com/wxlbd/gin-casbin-admin/internal/application/role"
	"github.com/wxlbd/gin-casbin-admin/internal/application/user"

	// "github.com/wxlbd/gin-casbin-admin/internal/handler" // Removing old handler import
	captchaInfra "github.com/wxlbd/gin-casbin-admin/internal/infrastructure/captcha"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/persistence"
	v1 "github.com/wxlbd/gin-casbin-admin/internal/interfaces/api/v1"

	// "github.com/wxlbd/gin-casbin-admin/internal/repository" // Removing old repository import
	"github.com/wxlbd/gin-casbin-admin/internal/server"
	// "github.com/wxlbd/gin-casbin-admin/internal/service" // Removing old service import
	"github.com/wxlbd/gin-casbin-admin/pkg/casbinx"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/gormx"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"github.com/wxlbd/gin-casbin-admin/pkg/redisx"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

// var RepositorySet = wire.NewSet(
// 	repository.NewRepository,
// )

// var ServiceSet = wire.NewSet(
// 	service.NewService,
// )

// var HandlerSet = wire.NewSet(
// 	handler.NewHandler,
// )

var UserSet = wire.NewSet(
	persistence.NewUserRepository,
	user.NewUserService,
	v1.NewUserHandler,
)

var RoleSet = wire.NewSet(
	persistence.NewRoleRepository,
	role.NewRoleService,
	v1.NewRoleHandler,
)

var MenuSet = wire.NewSet(
	persistence.NewMenuRepository,
	menu.NewMenuService,
	v1.NewMenuHandler,
)

var DictSet = wire.NewSet(
	persistence.NewDictRepository,
	dict.NewDictService,
	v1.NewDictHandler,
)

var CaptchaSet = wire.NewSet(
	captchaInfra.NewRedisStore,
	captcha.NewCaptchaService,
	v1.NewCaptchaHandler,
)

func NewWire(cfg *config.Config, logger *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		casbinx.New,
		gormx.NewDB,
		redisx.New,
		jwtx.New,
		ServerSet,
		// RepositorySet, // Removing old sets
		// ServiceSet,
		// HandlerSet,
		UserSet,
		RoleSet,
		MenuSet,
		DictSet,
		CaptchaSet,
	))

	return nil, nil, nil
}
