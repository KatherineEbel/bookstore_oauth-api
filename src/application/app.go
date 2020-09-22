package application

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/KatherineEbel/bookstore_oauth-api/src/domain/accessToken"
	h "github.com/KatherineEbel/bookstore_oauth-api/src/http"
	"github.com/KatherineEbel/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func Start() {
	tokenHandler := h.Handler(accessToken.Service(db.Repository()))
	router.GET("/oauth/accessToken/:id", tokenHandler.GetById)
	router.POST("/oauth/accessToken", tokenHandler.Create)
	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
