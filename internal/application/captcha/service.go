package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
	"github.com/wxlbd/gin-casbin-admin/internal/infrastructure/captcha"
)

type Service interface {
	Generate(ctx context.Context) (id, b64s string, err error)
	Verify(ctx context.Context, id, answer string) bool
}

type captchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService(store *captcha.RedisStore) Service {
	return &captchaService{
		store: store,
	}
}

func (s *captchaService) Generate(ctx context.Context) (id, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(180, 360, 4, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, _, err = c.Generate()
	return
}

func (s *captchaService) Verify(ctx context.Context, id, answer string) bool {
	return s.store.Verify(id, answer, true)
}
