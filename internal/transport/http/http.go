package http

import (
	"context"

	"github.com/gin-gonic/gin"
)

type service interface {
	Get(url string, ctx context.Context) (string, error)
	Add(url string, ctx context.Context) error
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

	r.GET("/v1/get", s.Get)
	r.GET("/v1/add", s.Add)

	return r.Run(adr)
}