package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/router"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/core"
	log "github.com/sirupsen/logrus"
	logger "github.com/stremovskyy/gin-request-logger"
)

func NewAndRun(config core.ServerConfig) {
	server := NewServer(config)
	server.Run()
}

func NewServer(config core.ServerConfig) *gin.Engine {
	gin.SetMode(mapTopGinMode(config.Mode))
	server := gin.Default()

	if config.ShowLog {
		log.SetLevel(log.TraceLevel)
		server.Use(logger.RequestLogger(true))
	}

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	groupV1 := server.Group("/api/v1")
	router.NewTypeRouter(groupV1)
	router.NewGitmojiRouter(groupV1)

	return server
}

func mapTopGinMode(mode core.ServerMode) string {
	switch mode {
	case core.ServerProdMode:
		return gin.ReleaseMode
	case core.ServerTestingMode:
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
