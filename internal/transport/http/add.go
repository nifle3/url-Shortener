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
	result, err := s.service.Add(val, c.Request.Context())

	if err != nil {
		c.String(http.StatusInternalServerError, "service error %s", err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"url": result})
}
