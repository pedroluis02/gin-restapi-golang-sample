package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/controller"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/data/repository"
)

func NewTypeRouter(rg *gin.RouterGroup) {
	controller := controller.NewTypeController(&repository.ConventionalTypeRepositoryImpl{})

	router := rg.Group("/types")
	router.GET("", controller.FindAll)
}
