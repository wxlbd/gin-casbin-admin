package middleware

import (
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/pkg/jwt"
	"gin-casbin-admin/pkg/log"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(enforcer *casbin.Enforcer, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, exists := ctx.Get("claims")
		if !exists {
			ctx.Abort()
			return
		}
		claims := v.(*jwt.MyCustomClaims)
		sub, err := claims.GetSubject()
		if err != nil {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrForbidden, nil)
			ctx.Abort()
			return
		}
		if ok, _ := enforcer.Enforce(sub, ctx.Request.URL.Path, ctx.Request.Method); !ok {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrForbidden, nil)
			ctx.Abort()
			return
		}
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}
