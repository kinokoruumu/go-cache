package cache

import (
	"github.com/kinokoruumu/go-cache/service"
	"github.com/kinokoruumu/go-cache/model"
)

var Article = articleCache{}

type articleCache struct {
	Articles []model.Article
}

func (a *articleCache) FindAll() []model.Article {
	if a.Articles != nil {
		return a.Articles
	}
	articles := service.Article.FindAll()
	a.Articles = articles
	return articles
}
