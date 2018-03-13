package service

import "github.com/kinokoruumu/go-cache/model"

var ArticleTag = articleTagService{}

type articleTagService struct {}

func (u *articleTagService) Store(articleTag model.ArticleTag) model.ArticleTag {
	db.Create(&articleTag)
	return articleTag
}

func (u *articleTagService) FindByArticleID(articleID uint) []model.ArticleTag {
	var articleTags []model.ArticleTag
	db.Find(&articleTags, "article_id=?", articleID)
	return articleTags
}
