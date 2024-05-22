package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/controller"
)

func NewTypeRouter(r *gin.Engine) {

	router := r.Group("/types")
	router.GET("", controller.FindAllTypes)
}
