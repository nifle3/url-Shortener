package redis

import (
	"context"
)

func (s *Storage) Add(shortener, url string, ctx context.Context) error {
	return s.client.Set(ctx, shortener, url, 0).Err()
}
