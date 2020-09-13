package graph

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wei840222/blog/db"
	"github.com/wei840222/blog/graph/model"
)

func graphQLIDToDBID(id string) uint {
	s := strings.Split(id, ":")
	i, _ := strconv.Atoi(s[1])
	return uint(i)
}

func dbUserToGraphQLUser(user *db.User) *model.User {
	return &model.User{
		ID:         fmt.Sprintf("user:%d", user.ID),
		Name:       user.Name,
		AvatarURL:  db.StringToStringPtr(user.AvatarURL.String),
		LineUserID: db.StringToStringPtr(user.LINEUserID.String),
	}
}

func dbPostToGraphQLPost(post *db.Post) *model.Post {
	return &model.Post{
		ID:   fmt.Sprintf("post:%d", post.ID),
		Text: post.Text,
	}
}

func dbCommentToGraphQLComment(comment *db.Comment) *model.Comment {
	return &model.Comment{
		ID:   fmt.Sprintf("comment:%d", comment.ID),
		Text: comment.Text,
	}
}
