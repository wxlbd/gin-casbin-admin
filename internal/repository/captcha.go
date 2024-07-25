package repository

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"time"
)

type captchaRepository struct {
	*Repository
}

func (c *captchaRepository) Set(id string, answer string) error {
	return c.rdb.Set(context.Background(), id, answer, time.Minute).Err()
}

func (c *captchaRepository) Get(id string, clear bool) string {
	var answer string
	answer = c.rdb.Get(context.Background(), id).String()
	if clear {
		_ = c.rdb.Del(context.Background(), id)
	}
	return answer
}

func (c *captchaRepository) Verify(id, answer string, clear bool) bool {
	stored := c.rdb.Get(context.Background(), id).String()
	if clear {
		_ = c.rdb.Del(context.Background(), id)
	}
	return stored == answer
}

func NewCaptchaRepository(
	r *Repository,
) base64Captcha.Store {
	return &captchaRepository{
		Repository: r,
	}
}
