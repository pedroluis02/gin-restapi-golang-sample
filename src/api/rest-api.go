package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/router"
	log "github.com/sirupsen/logrus"
	logger "github.com/stremovskyy/gin-request-logger"
)

func NewAndRun() {
	log.SetLevel(log.TraceLevel)

	server := gin.Default()
	server.Use(logger.RequestLogger(true))

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	groupV1 := server.Group("/api/v1")
	router.NewTypeRouter(groupV1)
	router.NewGitmojiRouter(groupV1)

	server.Run()
}
