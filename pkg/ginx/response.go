package ginx

import (
	"errors"

	"github.com/gin-gonic/gin"
	cerrors "github.com/wxlbd/gin-casbin-admin/pkg/errors"
)

const (
	SUCCESS = 200
	ERROR   = 500
)

type ListData struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code:    SUCCESS,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, businessCode int, message string, httpCode ...int) {
	hc := 200
	if len(httpCode) > 0 {
		hc = httpCode[0]
	}
	c.JSON(hc, Response{
		Code:    businessCode,
		Message: message,
	})
}

// ParamError 参数错误响应
func ParamError(c *gin.Context, err error) {
	_ = c.Error(err)
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, err error) {
	var customErr *cerrors.Error
	if errors.As(err, &customErr) {
		httpCode := 200
		if customErr.Status != 0 {
			httpCode = customErr.Status
		}
		Error(c, customErr.Code, customErr.Message, httpCode)
		return
	}
	Error(c, 401, "未授权", 401)
}

// Forbidden 禁止访问响应
func Forbidden(c *gin.Context) {
	Error(c, 403, "禁止访问", 403)
}

// ServerError 服务器错误响应
func ServerError(c *gin.Context, err error) {
	var customErr *cerrors.Error
	if errors.As(err, &customErr) {
		httpCode := 200
		if customErr.Status != 0 {
			httpCode = customErr.Status
		}
		Error(c, customErr.Code, customErr.Message, httpCode)
		return
	}
	Error(c, 500, err.Error(), 500)
}
