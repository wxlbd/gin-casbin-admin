package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/application/captcha"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type CaptchaHandler struct {
	svc captcha.Service
}

func NewCaptchaHandler(svc captcha.Service) *CaptchaHandler {
	return &CaptchaHandler{svc: svc}
}

func (h *CaptchaHandler) Generate(c *gin.Context) {
	id, b64s, err := h.svc.Generate(c)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, &captcha.CaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	})
}
