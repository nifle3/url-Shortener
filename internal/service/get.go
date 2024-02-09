package service

import "context"

func (u *UrlShortener) Get(url string, ctx context.Context) string {
	return u.database.Get(url, ctx)
}
