package inmemory

import "sync"

type Storage struct {
	mu      sync.Mutex
	storage map[string]string
}

func New() *Storage {
	return &Storage{
		mu:      sync.Mutex{},
		storage: make(map[string]string),
	}
}
