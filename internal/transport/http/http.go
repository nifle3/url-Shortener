package http

import (
	"context"

	"github.com/gin-gonic/gin"
)

type service interface {
	Get(url string, ctx context.Context) (string, error)
	Add(url string, ctx context.Context) (string, error)
}

type Server struct {
	service service
}

func New(service service) Server {
	return Server{
		service: service,
	}
}

func (s *Server) Listen(adr string) error {
	r := gin.Default()

	r.GET("/v1/get/:url", s.Get)
	r.GET("/v1/add/:url", s.Add)

	return r.Run(adr)
}
