package db_test

import (
	"github.com/wei840222/blog/db"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	db.TruncateAllTable()
	post, err := db.CreatePost(&db.Post{
		Text:   "TestCreatePost",
		UserID: 1,
	})

	assert.Nil(t, err)
	assert.Equal(t, "TestCreatePost", post.Text)
	assert.Equal(t, uint(1), post.UserID)
}

func TestListPosts(t *testing.T) {
	db.TruncateAllTable()
	_, err := db.CreatePost(&db.Post{
		Text:   "TestListPosts",
		UserID: 1,
	})
	assert.Nil(t, err)

	_, err = db.CreatePost(&db.Post{
		Text:   "TestListPosts2",
		UserID: 1,
	})
	assert.Nil(t, err)

	posts, err := db.ListPosts()

	assert.Nil(t, err)
	assert.Len(t, posts, 2)
}
