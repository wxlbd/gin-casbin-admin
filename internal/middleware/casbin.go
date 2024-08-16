package middleware

import (
	"fmt"
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/repository"
	"gin-casbin-admin/pkg/jwt"
	"gin-casbin-admin/pkg/log"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CasbinMiddleware(enforcer *casbin.Enforcer, menuRepo repository.MenuRepository, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, exists := ctx.Get("claims")
		if !exists {
			ctx.Abort()
			return
		}
		claims := v.(*jwt.MyCustomClaims)
		fmt.Printf("%+v\n", claims)
		sub, err := claims.GetSubject()
		if err != nil {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrForbidden, nil)
			ctx.Abort()
			return
		}
		fmt.Println(sub)
		menu, err := menuRepo.GetByPathMethod(ctx, ctx.Request.URL.Path, ctx.Request.Method)
		if err != nil {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrForbidden, nil)
			return
		}
		if ok, _ := enforcer.Enforce(sub, strconv.Itoa(int(menu.Id)), strconv.Itoa(int(menu.Type))); !ok {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrForbidden, nil)
			ctx.Abort()
			return
		}
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}
