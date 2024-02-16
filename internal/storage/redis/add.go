package redis

import (
	"context"
)

func (s *Storage) Add(shortener, url string, ctx context.Context) (string, error) {
	return s.client.Set(ctx, shortener, url, 0).Result()
}
