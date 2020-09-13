package db

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string
	PostID uint `gorm:"index"`
	UserID uint
}

func CreateComment(comment *Comment) (*Comment, error) {
	if err := db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func ListCommentsByPostID(postID uint) ([]*Comment, error) {
	var comments []*Comment
	if err := db.Where(Comment{
		PostID: postID,
	}).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
