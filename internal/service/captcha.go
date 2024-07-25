package service

import (
	"context"
	v1 "gin-casbin-admin/api/v1"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaService interface {
	GenerateCaptcha(ctx context.Context) (key string, b64s string, err error)
	VerifyCaptcha(ctx context.Context, captchaId, captchaValue string) error
}

type captchaService struct {
	*base64Captcha.Captcha
}

func (c *captchaService) GenerateCaptcha(ctx context.Context) (string, string, error) {
	key, b64s, answer, err := c.Generate()
	if err != nil {
		return "", "", err
	}
	if err := c.Store.Set(key, answer); err != nil {
		return "", "", err
	}
	return key, b64s, nil
}

func (c *captchaService) VerifyCaptcha(ctx context.Context, captchaId, captchaValue string) error {
	if !c.Verify(captchaId, captchaValue, true) {
		return v1.ErrCaptchaInvalid
	}
	return nil
}

func NewCaptchaService(store base64Captcha.Store) CaptchaService {
	driverString := base64Captcha.DriverString{
		Height:          50,
		Width:           110,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}

	captcha := base64Captcha.NewCaptcha(&driverString, store)
	return &captchaService{
		Captcha: captcha,
	}
}
