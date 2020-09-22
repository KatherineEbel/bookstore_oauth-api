package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	"github.com/KatherineEbel/bookstore_oauth-api/src/utils/errors"
)

type IHandler interface {
	GetById(*gin.Context)
	Create(ctx *gin.Context)
}

type handler struct {
	service accessToken.IService
}

func (h *handler) Create(c *gin.Context) {
	var t accessToken.AccessToken
	if err := c.ShouldBindJSON(&t); err != nil {
		e := errors.NewBadRequestError("invalid request")
		c.JSON(e.Code, e)
		return
	}
	fmt.Println(accessToken.GetNewAccessToken())
	rErr := h.service.Create(t)
	if rErr != nil {
		c.JSON(rErr.Code, rErr)
		return
	}
	c.JSON(http.StatusCreated, t)
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
