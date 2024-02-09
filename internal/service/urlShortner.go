package service

import "context"

type storage interface {
	Get(string, context.Context) string
	Add(string, string, context.Context)
}

type UrlShortener struct {
	database storage
}

func New(storage storage) UrlShortener {
	return UrlShortener{
		database: storage,
	}
}
