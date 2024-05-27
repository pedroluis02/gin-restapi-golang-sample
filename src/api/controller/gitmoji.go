package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/api/dto"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/common"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/domain/model"
	"github.com/pedroluis02/gin-restapi-golang-sample/src/domain/repository"
)

type GitmojiController struct {
	repository repository.GitmojiRepository
}

func NewGitmojiController(repository repository.GitmojiRepository) *GitmojiController {
	return &GitmojiController{repository: repository}
}

func (cc *GitmojiController) FindAll(c *gin.Context) {
	types := cc.repository.FindAll()
	dto := common.MapArrayWithFunc(types, cc.mapToDto)

	c.JSON(http.StatusOK, dto)
}

func (cc *GitmojiController) mapToDto(model model.Gitmoji) dto.GitmojiDto {
	return dto.GitmojiDto{Id: model.Id, Code: model.Code, Emoji: model.Emoji}
}
