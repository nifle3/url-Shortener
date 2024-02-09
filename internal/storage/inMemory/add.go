package inmemory

import (
	"context"
	"fmt"
)

func (s *Storage) Add(shortener, url string, _ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.storage[shortener]
	if !ok {
		return fmt.Errorf("already exis")
	}

	s.storage[shortener] = val
	return nil
}
