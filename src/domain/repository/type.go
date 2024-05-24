package repository

import "github.com/pedroluis02/gin-restapi-golang-sample/src/domain/model"

type ConventionalTypeRepository interface {
	FindAll() []model.ConventionalType
}
