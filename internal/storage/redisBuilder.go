package storage

import (
	"fmt"
	"strings"
	"urlShortener/internal/storage/redis"
)

type RedisBuilder struct {
	addr, password string
	db             int
}

func NewRedisBuilder() RedisBuilder {
	return RedisBuilder{
		addr:     "localhost:6379",
		password: "",
		db:       0,
	}
}

func (r *RedisBuilder) Build() *redis.Storage {
	return redis.New(r.addr, r.password, r.db)
}

func (r *RedisBuilder) SetAddr(addr string) error {
	if len(addr) == 0 {
		return fmt.Errorf("addr cannot be empty")
	}

	if !strings.Contains(addr, ":") {
		return fmt.Errorf("addr must be in the format of host:port")
	}

	r.addr = addr
	return nil
}

func (r *RedisBuilder) SetPassword(password string) {
	r.password = password
}

func (r *RedisBuilder) SetDB(db int) error {
	if db < 0 {
		return fmt.Errorf("db cannot be less than 0")
	}

	r.db = db

	return nil
}
