package repository

import "github.com/pedroluis02/gin-restapi-golang-sample/src/domain/model"

var types = []model.ConventionalType{
	{Id: 1, Name: "feat"},
	{Id: 2, Name: "fix"},
	{Id: 3, Name: "refactor"},
	{Id: 4, Name: "build"},
	{Id: 5, Name: "chore"},
}

type ConventionalTypeRepositoryImpl struct {
}

func (r *ConventionalTypeRepositoryImpl) FindAll() []model.ConventionalType {
	return types
}
