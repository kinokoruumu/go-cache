package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Token struct {
	Model
	UserID uint
	Body   string
}

type Recipi struct {
	Model
}

type Article struct {
	Model
	Title string
	Text string
	Tags []Tag `gorm:"-"`
}

type Tag struct {
	Model
	Name string
}

type ArticleTag struct {
	Model
	ArticleID uint
	TagID uint
}
