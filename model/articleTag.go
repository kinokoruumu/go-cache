package model

func(self ArticleTag) GetTag() *Tag {
	var tag []Tag
	db.Find(&tag, "id=?", self.TagID)
	if len(tag) == 0 {
		return nil
	}
	// TODO panicでいいかまっきーに聞く
	return &tag[0]
}
