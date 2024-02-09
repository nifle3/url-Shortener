package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Add(c *gin.Context) {
	val, ok := c.Params.Get("url")
	if !ok {
		c.String(http.StatusBadRequest, "Parametr url is not exist")
		return
	}

	if err := s.service.Add(val, c.Request.Context()); err != nil {
		c.String(http.StatusInternalServerError, "service error %s", err.Error())
		return
	}

	c.String(http.StatusCreated, "url shortener created")
}
