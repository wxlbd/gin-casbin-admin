package server

import (
	apiV1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/docs"
	"gin-casbin-admin/internal/handler"
	"gin-casbin-admin/internal/middleware"
	"gin-casbin-admin/internal/repository"
	"gin-casbin-admin/pkg/jwt"
	"gin-casbin-admin/pkg/log"
	"gin-casbin-admin/pkg/server/http"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	enforcer *casbin.Enforcer,
	userHandler *handler.AdminUserHandler,
	roleHandler *handler.RoleHandler,
	captchaHandler *handler.CaptchaHandler,
	menuHandler *handler.MenuHandler,
	menuRepo repository.MenuRepository,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/api/v1")
	{
		v1.GET("/captcha", captchaHandler.GetCaptcha)
		// No route group has permission
		noAuthRouter := v1.Group("/user")
		{
			noAuthRouter.POST("/register", userHandler.AddAdminUser)
			noAuthRouter.POST("/login", userHandler.Login)
			noAuthRouter.POST("refreshToken", userHandler.RefreshToken)
		}
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.GET("/user", userHandler.GetProfile)
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
			strictAuthRouter.POST("/user/role", userHandler.SetUserRoles)
		}

		menuRouter := v1.Group("/menu")
		{
			menuRouter.POST("", menuHandler.AddMenu)
			menuRouter.GET("", menuHandler.GetMenuList)
			menuRouter.PUT("/:id", menuHandler.UpdateMenu)
			menuRouter.DELETE("", menuHandler.DeleteMenu)
			menuRouter.GET("/tree", menuHandler.GetMenuTree)
			menuRouter.GET("/:id", menuHandler.GetMenu)
		}

		roleRouter := v1.Group("/role").Use(middleware.StrictAuth(jwt, logger), middleware.CasbinMiddleware(enforcer, menuRepo, logger))
		{
			roleRouter.POST("", roleHandler.Add)
			roleRouter.GET("/:id", roleHandler.Get)
			roleRouter.GET("", roleHandler.List)
			roleRouter.PUT("/:id", roleHandler.Update)
			roleRouter.DELETE("/:id", roleHandler.Delete)
		}
	}

	return s
}
