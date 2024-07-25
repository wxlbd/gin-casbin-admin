package handler

import (
	v1 "gin-casbin-admin/api/v1"
	"gin-casbin-admin/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaHandler struct {
	*Handler
	service.CaptchaService
}

func NewCaptchaHandler(handler *Handler, captchaService service.CaptchaService) *CaptchaHandler {
	return &CaptchaHandler{
		Handler:        handler,
		CaptchaService: captchaService,
	}
}

func (c *CaptchaHandler) GetCaptcha(ctx *gin.Context) {
	key, b64s, err := c.CaptchaService.GenerateCaptcha(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusOK, err, nil)
		return
	}

	v1.HandleSuccess(ctx, &v1.CaptchaResponseData{
		Key:  key,
		B64s: b64s,
	})
}
