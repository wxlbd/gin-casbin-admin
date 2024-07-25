//go:build wireinject
// +build wireinject

package wire

import (
	"gin-casbin-admin/internal/handler"
	"gin-casbin-admin/internal/repository"
	"gin-casbin-admin/internal/server"
	"gin-casbin-admin/internal/service"
	"gin-casbin-admin/pkg/app"
	"gin-casbin-admin/pkg/casbin"
	"gin-casbin-admin/pkg/jwt"
	"gin-casbin-admin/pkg/log"
	"gin-casbin-admin/pkg/server/http"
	"gin-casbin-admin/pkg/sid"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewPermissionsRepository,
	repository.NewRoleRepository,
	repository.NewCaptchaRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewPermissionsService,
	service.NewRoleService,
	service.NewCaptchaService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewPermissionHandler,
	handler.NewRoleHandler,
	handler.NewCaptchaHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		casbin.NewCasbinEnforcer,
		newApp,
	))
}
