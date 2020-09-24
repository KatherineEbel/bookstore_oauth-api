package application

import (
	"log"

	"github.com/gin-gonic/gin"

	h "github.com/KatherineEbel/bookstore_oauth-api/src/http"
	"github.com/KatherineEbel/bookstore_oauth-api/src/repository/db"
	"github.com/KatherineEbel/bookstore_oauth-api/src/repository/rest"
	accessToken2 "github.com/KatherineEbel/bookstore_oauth-api/src/services/accessToken"
)

var (
	router = gin.Default()
)

func Start() {
	tokenHandler := h.Handler(accessToken2.Service(db.Repository(), rest.Repository()))
	router.GET("/oauth/access_token/:id", tokenHandler.GetById)
	router.POST("/oauth/access_token", tokenHandler.Create)
	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
