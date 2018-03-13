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
}

