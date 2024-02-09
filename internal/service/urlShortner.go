package service

import "context"

type storage interface {
	Get(string, context.Context) (string, error)
	Add(string, string, context.Context) error
}

type UrlShortener struct {
	database storage
}

func New(storage storage) UrlShortener {
	return UrlShortener{
		database: storage,
	}
}
