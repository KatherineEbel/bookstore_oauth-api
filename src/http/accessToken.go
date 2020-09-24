package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	accessToken2 "github.com/KatherineEbel/bookstore_oauth-api/src/services/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

type IHandler interface {
	GetById(*gin.Context)
	Create(ctx *gin.Context)
}

type handler struct {
	service accessToken2.IService
}

func Handler(service accessToken2.IService) IHandler {
	return &handler{service}
}

func (h *handler) Create(c *gin.Context) {
	var r accessToken.Request
	if err := c.ShouldBindJSON(&r); err != nil {
		e := errors.NewBadRequestError("invalid request")
		c.JSON(e.Code, e)
		return
	}
	t, rErr := h.service.Create(r)
	if rErr != nil {
		c.JSON(rErr.Code, rErr)
		return
	}
	c.JSON(http.StatusCreated, t)
}

func (h *handler) GetById(c *gin.Context) {
	t, err := h.service.GetBId(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, t)
}
