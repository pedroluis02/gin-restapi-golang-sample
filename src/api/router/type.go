package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/controller"
)

func NewTypeRouter(rg *gin.RouterGroup) {
	controller := controller.ConventionalTypeController{}

	router := rg.Group("/types")
	router.GET("", controller.FindAll)
}
