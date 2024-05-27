package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/controller"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/data/repository"
)

func NewGitmojiRouter(rg *gin.RouterGroup) {
	controller := controller.NewGitmojiController(&repository.GitmojiRepositoryImpl{})

	router := rg.Group("/gitmojis")
	router.GET("", controller.FindAll)
}
