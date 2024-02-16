package redis

import "context"

func (s *Storage) Get(shortener string, _ context.Context) (string, error) {

	return "val", nil
}
