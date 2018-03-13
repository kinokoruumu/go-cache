package main

import (
	"github.com/kinokoruumu/go-cache/service"
	"github.com/kinokoruumu/go-cache/model"
)

func main() {
	service.Article.Store(model.Article{
		Title: "タイトル1",
		Text: "これは本文です。",
	})
	service.Article.Store(model.Article{
		Title: "タイトル2",
		Text: "これは本文です。",
	})
	service.Article.Store(model.Article{
		Title: "タイトル3",
		Text: "これは本文です。",
	})


	service.Tag.Store(model.Tag{
		Name: "tag1",
	})
	service.Tag.Store(model.Tag{
		Name: "tag2",
	})
	service.Tag.Store(model.Tag{
		Name: "tag3",
	})
	service.Tag.Store(model.Tag{
		Name: "tag4",
	})
	service.Tag.Store(model.Tag{
		Name: "tag5",
	})

	service.ArticleTag.Store(model.ArticleTag{
		ArticleID: 1,
		TagID: 1,
	})
	service.ArticleTag.Store(model.ArticleTag{
		ArticleID: 1,
		TagID: 2,
	})
	service.ArticleTag.Store(model.ArticleTag{
		ArticleID: 1,
		TagID: 3,
	})

	service.ArticleTag.Store(model.ArticleTag{
		ArticleID: 2,
		TagID: 1,
	})
	service.ArticleTag.Store(model.ArticleTag{
		ArticleID: 2,
		TagID: 4,
	})
}

