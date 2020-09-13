package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wei840222/blog/internal/db"
	"github.com/wei840222/blog/internal/graph/generated"
	"github.com/wei840222/blog/internal/graph/model"
)

func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	user, err := db.GetUser(graphQLIDToDBID(obj.User.ID))
	if err != nil {
		return nil, err
	}
	return dbUserToGraphQLUser(user), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := db.CreateUser(&db.User{
		Name:       input.Name,
		AvatarURL:  db.StringPtrToSqlNullString(input.AvatarURL),
		LINEUserID: db.StringPtrToSqlNullString(input.LineUserID),
	})
	if err != nil {
		return nil, err
	}
	return dbUserToGraphQLUser(user), nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post, err := db.CreatePost(&db.Post{
		Text:   input.Text,
		UserID: graphQLIDToDBID(input.UserID),
	})
	if err != nil {
		return nil, err
	}
	return dbPostToGraphQLPost(post), nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	comment, err := db.CreateComment(&db.Comment{
		Text:   input.Text,
		PostID: graphQLIDToDBID(input.PostID),
		UserID: graphQLIDToDBID(input.UserID),
	})
	if err != nil {
		return nil, err
	}
	return dbCommentToGraphQLComment(comment), nil
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	user, err := db.GetUser(graphQLIDToDBID(obj.User.ID))
	if err != nil {
		return nil, err
	}
	return dbUserToGraphQLUser(user), nil
}

func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	comments, err := db.ListCommentsByPostID(graphQLIDToDBID(obj.ID))
	if err != nil {
		return nil, err
	}
	var gqlComments []*model.Comment
	for _, comment := range comments {
		gqlComments = append(gqlComments, dbCommentToGraphQLComment(comment))
	}
	return gqlComments, nil
}

func (r *queryResolver) LineUser(ctx context.Context, lineUserID string) (*model.User, error) {
	user, err := db.GetUserByLINEUserID(lineUserID)
	if err != nil {
		return nil, err
	}
	return dbUserToGraphQLUser(user), nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := db.ListPosts()
	if err != nil {
		return nil, err
	}
	var gqlPosts []*model.Post
	for _, post := range posts {
		gqlPosts = append(gqlPosts, dbPostToGraphQLPost(post))
	}
	return gqlPosts, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
