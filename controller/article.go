package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kinokoruumu/go-cache/model"
	"github.com/kinokoruumu/go-cache/service"
	"github.com/kinokoruumu/go-cache/cache"
)

var Article = articleController{}

type articleController struct {
}

func (a *articleController) Create(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		BatRequest(err.Error(), c)
		return
	}
	article = service.Article.Store(article)
	Json(article, c)
}

func (a *articleController) FindAll(c *gin.Context) {
	articles := cache.Article.FindAll()
	Json(articles, c)
}
