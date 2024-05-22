package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/router"
)

func main() {
	log.Println("GIN Restful API")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.NewTypeRouter(r)

	r.Run()
}
