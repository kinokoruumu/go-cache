package service

import (
	"github.com/kinokoruumu/go-cache/model"
)

var Article = article{}

type article struct {}

func (u *article) Store(article model.Article) model.Article {
	db.Create(&article)
	return article
}

func (u *article) FindAll() []model.Article {
	var articles []model.Article
	db.Find(&articles)
	var resArticles []model.Article
	for _, v := range articles {
		articleTag := ArticleTag.FindByArticleID(v.ID)
		for _, vv := range articleTag {
			tag := vv.GetTag()
			v.Tags = append(v.Tags, *tag)
		}
		resArticles = append(resArticles, v)
	}
	return resArticles
}
