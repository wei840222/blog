package db

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text   string
	UserID uint `gorm:"index"`
}

func CreatePost(post *Post) (*Post, error) {
	if err := db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func ListPosts() ([]*Post, error) {
	var posts []*Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
