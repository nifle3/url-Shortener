package storage

import "urlShortener/internal/storage/redis"

type RedisBuilder struct {
	addr, password string
	db             int
}

func NewRedisBuilder() RedisBuilder {
	defaultValue := RedisBuilder{}
	defaultValue.SetPassword("")
	defaultValue.SetDB(0)
	defaultValue.SetAddr("localhost:6379")

	return defaultValue
}

func (r *RedisBuilder) Build() *redis.Storage {
	return redis.New(r.addr, r.password, r.db)
}

func (r *RedisBuilder) SetAddr(addr string) {
	if len(addr) == 0 {
		return
	}

	r.addr = addr
}

func (r *RedisBuilder) SetPassword(password string) {
	r.password = password
}

func (r *RedisBuilder) SetDB(db int) {
	if db > 0 {
		db = 0
	}

	r.db = db
}
