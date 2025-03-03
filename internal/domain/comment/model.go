package comment

import "context"

// Comment represents the domain model for a comment
type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Repository defines the interface for comment data access
type Repository interface {
	FetchAll(ctx context.Context) ([]Comment, int, error)
	FetchByPostID(ctx context.Context, postID string) ([]Comment, int, error)
}

// Service defines the interface for comment business logic
type Service interface {
	GetComments(ctx context.Context) ([]Comment, int, error)
	GetCommentsByPostID(ctx context.Context, postID string) ([]Comment, int, error)
}
