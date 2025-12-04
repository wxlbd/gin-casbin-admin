package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wxlbd/gin-casbin-admin/internal/application/user"
	v1 "github.com/wxlbd/gin-casbin-admin/internal/interfaces/api/v1"
	"github.com/wxlbd/gin-casbin-admin/internal/middleware"
	"github.com/wxlbd/gin-casbin-admin/pkg/config"
	"github.com/wxlbd/gin-casbin-admin/pkg/jwtx"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

func NewServerHTTP(
	cfg *config.Config,
	logger *log.Logger,
	jwt *jwtx.JWT,
	userHandler *v1.UserHandler,
	roleHandler *v1.RoleHandler,
	menuHandler *v1.MenuHandler,
	dictHandler *v1.DictHandler,
	captchaHandler *v1.CaptchaHandler,
	enforcer *casbin.Enforcer,
	userService user.Service,
) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default() // 注册中间件
	r.Use(middleware.RequestLogger(logger))
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	// 注册路由
	api := r.Group("/api/v1")
	{
		// 认证相关接口
		auth := api.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh-token", userHandler.RefreshToken)
			auth.POST("/logout", userHandler.Logout)
			auth.GET("/captcha", captchaHandler.Generate)

			// 需要JWT认证的接口
			jwtGroup := api.Group("")
			jwtGroup.Use(middleware.JWTAuth(jwt))
			{
				profile := jwtGroup.Group("user/profile")
				{
					profile.GET("", userHandler.Current)
					profile.GET("menus", menuHandler.GetUserMenuTree)
					profile.GET("roles", roleHandler.GetPermittedMenus)
				}

				// 需要完整权限控制的接口
				sys := jwtGroup.Group("/system")
				sys.Use(middleware.CasbinMiddleware(enforcer, logger, userService))

				jwtGroup.GET("system/role/all", roleHandler.GetAllRoles)
			}

			// 需要完整权限控制的接口
			authorized := api.Group("")
			authorized.Use(
				middleware.JWTAuth(jwt),
				middleware.CasbinMiddleware(enforcer, logger, userService),
			)
			sys := authorized.Group("system")

			// 权限控制
			{
				// 用户管理 system:user:xxx
				userGroup := sys.Group("user")
				{
					userGroup.GET("", userHandler.List)                              // system:user:list
					userGroup.POST("", userHandler.Create)                           // system:user:create
					userGroup.PUT("/:id", userHandler.Update)                        // system:user:update
					userGroup.DELETE("/:ids", userHandler.Delete)                    // system:user:delete
					userGroup.GET("/:id", userHandler.Detail)                        // system:user:detail
					userGroup.GET("/:id/roles", userHandler.GetUserRoles)            // system:user:get:roles
					userGroup.PUT("/:id/password", userHandler.ResetPassword)        // system:user:set:password
					userGroup.PUT("/:id/roles", userHandler.AssignRoles)             // system:user:set:roles
					userGroup.GET("/current/roles", userHandler.GetCurrentUserRoles) // 获取当前用户角色
				}

				// 角色管理 permission:role:xxx
				roleGroup := sys.Group("role")
				{
					roleGroup.GET("", roleHandler.List)                            // permission:role:list
					roleGroup.POST("", roleHandler.Create)                         // permission:role:create
					roleGroup.PUT("/:id", roleHandler.Update)                      // permission:role:update
					roleGroup.DELETE("/:ids", roleHandler.Delete)                  // permission:role:delete
					roleGroup.GET("/:id", roleHandler.Detail)                      // permission:role:detail
					roleGroup.GET("/:id/menus", roleHandler.GetPermittedMenus)     // permission:role:menu:list
					roleGroup.POST("/:id/menus", roleHandler.AssignRoleMenusByIDs) // permission:role:menu:assignus
				}

				// 菜单管理 permission:menu:xxx
				menuGroup := sys.Group("menu")
				{
					menuGroup.GET("", menuHandler.List)                      // system:menu:list
					menuGroup.POST("", menuHandler.Create)                   // system:menu:create
					menuGroup.PUT("/:id", menuHandler.Update)                // system:menu:update
					menuGroup.DELETE("/:ids", menuHandler.Delete)            // system:menu:delete
					menuGroup.GET("/tree", menuHandler.GetMenuTree)          // system:menu:tree
					menuGroup.GET("/user-tree", menuHandler.GetUserMenuTree) // system:menu:user-tree
				}

				// 字典管理
				{
					// 字典类型管理
					dictType := sys.Group("dict-type")
					{
						dictType.GET("", dictHandler.ListDictType)           // system:dict:type:list
						dictType.POST("", dictHandler.CreateDictType)        // system:dict:type:create
						dictType.PUT("/:id", dictHandler.UpdateDictType)     // system:dict:type:update
						dictType.DELETE("/:ids", dictHandler.DeleteDictType) // system:dict:type:delete
						dictType.GET("/:id", dictHandler.GetDictType)        // system:dict:type:detail
					}

					// 字典数据管理
					dictData := sys.Group("dict-data")
					{
						dictData.GET("", dictHandler.ListDictData)                 // system:dict:data:list
						dictData.POST("", dictHandler.CreateDictData)              // system:dict:data:create
						dictData.PUT("/:id", dictHandler.UpdateDictData)           // system:dict:data:update
						dictData.DELETE("/:ids", dictHandler.DeleteDictData)       // system:dict:data:delete
						dictData.GET("/:id", dictHandler.GetDictData)              // system:dict:data:detail
						dictData.GET("/type/:type", dictHandler.GetDictDataByType) // system:dict:data:typelist:type
					}
				}
			}
		}

	}

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
