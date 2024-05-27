package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/dto"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/common"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/domain/model"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/domain/repository"
)

type ConventionalTypeController struct {
	repository repository.ConventionalTypeRepository
}

func NewTypeController(repository repository.ConventionalTypeRepository) *ConventionalTypeController {
	return &ConventionalTypeController{repository: repository}
}

func (cc *ConventionalTypeController) FindAll(c *gin.Context) {
	types := cc.repository.FindAll()
	dto := common.MapArrayWithFunc(types, cc.mapToDto)

	c.JSON(http.StatusOK, dto)
}

func (cc *ConventionalTypeController) mapToDto(model model.ConventionalType) dto.ConventionalTypeDto {
	return dto.ConventionalTypeDto{Id: model.Id, Name: model.Name}
}
