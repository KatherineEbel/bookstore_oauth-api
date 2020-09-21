package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
)

type IHandler interface {
	GetById(*gin.Context)
}

type handler struct {
	service accessToken.IService
}

func Handler(service accessToken.IService) IHandler {
	return &handler{service}
}

func (h *handler) GetById(c *gin.Context) {
	t, err := h.service.GetBId(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, t)
}
