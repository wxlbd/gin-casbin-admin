package middleware

import (
	"net/http"

	"github.com/wxlbd/gin-casbin-admin/internal/application/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
	"go.uber.org/zap"
)

func CasbinMiddleware(enforcer *casbin.Enforcer, log *log.Logger, svc user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前登录用户
		userID := c.GetUint64("user_id")
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未登录",
			})
			c.Abort()
			return
		}

		// 获取用户的角色列表
		roles, err := svc.GetUserRoles(c, userID)
		if err != nil {
			log.WithContext(c).Error("获取用户角色失败", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取用户角色失败",
			})
			c.Abort()
			return
		}

		// 将角色列表存入上下文
		c.Set("user_roles", roles)

		// 检查是否是超级管理员
		isAdmin := false
		for _, role := range roles {
			if role.Code == "SuperAdmin" {
				isAdmin = true
				break
			}
		}

		// 如果是超级管理员，直接放行
		if isAdmin {
			c.Next()
			return
		}

		// 获取请求的URI和方法
		obj := c.Request.URL.Path
		act := c.Request.Method
		// 检查权限
		// 遍历用户角色，检查是否有权限
		hasPermission := false
		for _, role := range roles {
			ok, err := enforcer.Enforce(role.Code, obj, act)
			if err != nil {
				log.WithContext(c).Error("权限检查失败", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "权限检查失败",
				})
				c.Abort()
				return
			}
			if ok {
				hasPermission = true
				break
			}
		}
		// 如果没有任何角色有权限，返回403
		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "没有权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
