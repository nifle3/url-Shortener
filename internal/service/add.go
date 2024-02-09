package service

import (
	"context"
	"crypto/sha256"
)

func (u *UrlShortener) Add(url string, ctx context.Context) {
	hashed := sha256.New().Sum([]byte(url))
	u.database.Add(string(hashed), url, ctx)
}
