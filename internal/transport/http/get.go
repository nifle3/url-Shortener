package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Get(c *gin.Context) {
	val, ok := c.Params.Get("url")
	if !ok {
		c.String(http.StatusBadRequest, "parametr url is not set")
		return
	}

	urlRedirect, err := s.service.Get(val, c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "service error %s", err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, urlRedirect)
}
