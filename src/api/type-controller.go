package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func findAllTypes(c *gin.Context) {
	var types = []ConventionalTypeDto{
		{Id: 1, Name: "feat"},
		{Id: 2, Name: "fix"},
		{Id: 3, Name: "refactor"},
		{Id: 4, Name: "build"},
		{Id: 5, Name: "chore"},
	}

	c.JSON(http.StatusOK, types)
}
