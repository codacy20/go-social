package comment

import (
	"context"
	"social/internal/domain/comment"

	log "github.com/sirupsen/logrus"
)

// CommentService implements the comment.Service interface
type CommentService struct {
	repo   comment.Repository
	logger *log.Logger
}

// NewCommentService creates a new instance of CommentService
func NewCommentService(repo comment.Repository, logger *log.Logger) *CommentService {
	return &CommentService{
		repo:   repo,
		logger: logger,
	}
}

// GetComments retrieves all comments
func (s *CommentService) GetComments(ctx context.Context) ([]comment.Comment, int, error) {
	s.logger.Info("Service: Fetching all comments")
	return s.repo.FetchAll(ctx)
}

// GetCommentsByPostID retrieves comments for a specific post
func (s *CommentService) GetCommentsByPostID(ctx context.Context, postID string) ([]comment.Comment, int, error) {
	s.logger.Infof("Service: Fetching comments for post %s", postID)
	return s.repo.FetchByPostID(ctx, postID)
}
