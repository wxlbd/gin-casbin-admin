package repository

import (
	"context"
	"fmt"
	"time"
)

type TokenRepository interface {
	StoreRefreshToken(ctx context.Context, userId string, refreshToken string, expire time.Duration) error
	GetRefreshToken(ctx context.Context, userId string) (string, error)
	RevokeRefreshToken(ctx context.Context, userId string) error
}

const RefreshTokenKeyTemplate = "refresh_token:%s"

type tokenRepository struct {
	*Repository
}

func (t *tokenRepository) StoreRefreshToken(ctx context.Context, userId string, refreshToken string, expire time.Duration) error {
	return t.rdb.Set(ctx, fmt.Sprintf(RefreshTokenKeyTemplate, userId), refreshToken, expire).Err()
}

func (t *tokenRepository) GetRefreshToken(ctx context.Context, userId string) (string, error) {
	return t.rdb.Get(ctx, fmt.Sprintf(RefreshTokenKeyTemplate, userId)).Result()
}

func (t *tokenRepository) RevokeRefreshToken(ctx context.Context, userId string) error {
	return t.rdb.Del(ctx, fmt.Sprintf(RefreshTokenKeyTemplate, userId)).Err()
}

func NewTokenRepository(
	r *Repository,
) TokenRepository {
	return &tokenRepository{
		Repository: r,
	}
}
