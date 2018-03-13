package main

import "github.com/kinokoruumu/go-cache/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Token{})
	db.DropTableIfExists(&model.Recipi{})
	db.DropTableIfExists(&model.Article{})
	db.DropTableIfExists(&model.Tag{})
	db.DropTableIfExists(&model.ArticleTag{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Token{})
	db.AutoMigrate(&model.Recipi{})
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.ArticleTag{})
}
