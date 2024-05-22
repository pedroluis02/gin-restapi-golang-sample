package api

import "github.com/gin-gonic/gin"

func NewTypeRouter(r *gin.Engine) {

	router := r.Group("/types")
	router.GET("", findAllTypes)
}
