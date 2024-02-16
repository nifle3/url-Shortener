package redis

import "github.com/redis/go-redis/v9"

type Storage struct {
	client *redis.Client
}

func New(addr, password string, db int) *Storage {
	return &Storage{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}
