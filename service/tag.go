package service

import "github.com/kinokoruumu/go-cache/model"

var Tag = tag{}

type tag struct {}

func (u *tag) Store(tag model.Tag) model.Tag {
	db.Create(&tag)
	return tag
}

func (u *tag) Find(tagID uint) *model.Tag {
	var tags []model.Tag
	db.Find(&tags, "id=?", tagID)
	if len(tags) == 0 {
		return nil
	}
	return &tags[0]
}
