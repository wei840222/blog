package db_test

import (
	"github.com/wei840222/blog/db"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	db.TruncateAllTable()
	comment, err := db.CreateComment(&db.Comment{
		Text:   "TestCreateComment",
		PostID: 1,
		UserID: 1,
	})

	assert.Nil(t, err)
	assert.Equal(t, "TestCreateComment", comment.Text)
	assert.Equal(t, uint(1), comment.PostID)
	assert.Equal(t, uint(1), comment.UserID)
}

func TestListCommentsByPostID(t *testing.T) {
	db.TruncateAllTable()
	_, err := db.CreateComment(&db.Comment{
		Text:   "TestListCommentsByPostID",
		PostID: 1,
		UserID: 1,
	})
	assert.Nil(t, err)
	_, err = db.CreateComment(&db.Comment{
		Text:   "TestListCommentsByPostID2",
		PostID: 1,
		UserID: 1,
	})
	assert.Nil(t, err)

	comments, err := db.ListCommentsByPostID(1)
	assert.Nil(t, err)
	assert.Len(t, comments, 2)
}
