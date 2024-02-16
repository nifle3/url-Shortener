package redis

import "context"

func (s *Storage) Get(shortener string, ctx context.Context) (string, error) {
	return s.client.Get(ctx, shortener).Result()
}
