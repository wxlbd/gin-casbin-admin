package captcha

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client     *redis.Client
	expiration time.Duration
	keyPrefix  string
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{
		client:     client,
		expiration: 10 * time.Minute,
		keyPrefix:  "captcha:",
	}
}

func (s *RedisStore) Set(id string, value string) error {
	ctx := context.Background()
	return s.client.Set(ctx, s.keyPrefix+id, value, s.expiration).Err()
}

func (s *RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	key := s.keyPrefix + id
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(ctx, key)
	}
	return val
}

func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
