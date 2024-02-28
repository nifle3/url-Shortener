package http

import (
	"context"
	"fmt"
	"net/http"
)

type service interface {
	Get(url string, ctx context.Context) (string, error)
	Add(url string, ctx context.Context) (string, error)
}

type Server struct {
	service service
}

func New(service service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Listen(adr string) error {
	r := http.NewServeMux()

	r.HandleFunc(fmt.Sprintf("%s /v1/add", http.MethodPost), s.Add)
	r.HandleFunc(fmt.Sprintf("%s /v1/get/{url}", http.MethodGet), s.Get)

	return http.ListenAndServe(adr, r)
}
