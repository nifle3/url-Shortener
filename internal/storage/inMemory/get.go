package inmemory

import (
	"context"
	"fmt"
)

func (s *Storage) Get(shortener string, _ context.Context) (string, error) {
	val, ok := s.storage[shortener]
	if !ok {
		return "", fmt.Errorf("doesn't exist")
	}

	return val, nil
}
