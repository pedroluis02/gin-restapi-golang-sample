package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/dto"
)

func FindAllTypes(c *gin.Context) {
	var types = []dto.ConventionalTypeDto{
		{Id: 1, Name: "feat"},
		{Id: 2, Name: "fix"},
		{Id: 3, Name: "refactor"},
		{Id: 4, Name: "build"},
		{Id: 5, Name: "chore"},
	}

	c.JSON(http.StatusOK, types)
}
