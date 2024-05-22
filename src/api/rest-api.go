package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/router"
)

func NewAndRun() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	groupV1 := server.Group("/api/v1")

	router.NewTypeRouter(groupV1)

	server.Run()
}
