package test

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/core"
)

func NewServerTesting() *gin.Engine {
	config := core.ServerConfig{
		Mode:    core.ServerTestingMode,
		ShowLog: false,
	}
	return api.NewServer(config)
}
