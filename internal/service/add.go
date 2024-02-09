package service

import (
	"context"
	"crypto/sha256"
)

func (u *UrlShortener) Add(url string, ctx context.Context) error {
	hashed := sha256.New().Sum([]byte(url))
	return u.database.Add(string(hashed), url, ctx)
}
