package repository

import "github.com/pedroluis02/gin-restapi-golang-sample/src/domain/model"

var gitmojis = []model.Gitmoji{
	{Id: 1, Code: ":sparkles:", Emoji: "✨"},
	{Id: 2, Code: ":bug:", Emoji: "🐛"},
	{Id: 3, Code: ":hammer:", Emoji: "🔨"},
	{Id: 4, Code: ":package:", Emoji: "📦️"},
	{Id: 5, Code: ":wrench:", Emoji: "🔧"},
}

type GitmojiRepositoryImpl struct {
}

func (r *GitmojiRepositoryImpl) FindAll() []model.Gitmoji {
	return gitmojis
}
