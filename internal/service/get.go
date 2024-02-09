package service

import "context"

func (u *UrlShortener) Get(url string, ctx context.Context) (string, error) {
	return u.database.Get(url, ctx)
}
