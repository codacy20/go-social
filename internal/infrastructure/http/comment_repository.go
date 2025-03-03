package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"social/internal/domain/comment"
	"time"

	validator "social/depx/validation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	log "github.com/sirupsen/logrus"
)

// CommentRepository implements the comment.Repository interface
type CommentRepository struct {
	client  *http.Client
	baseURL string
	logger  *log.Logger
}

// NewCommentRepository creates a new instance of CommentRepository
func NewCommentRepository(logger *log.Logger) *CommentRepository {
	return &CommentRepository{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://jsonplaceholder.typicode.com",
		logger:  logger,
	}
}

func (r *CommentRepository) FetchAll(ctx context.Context) ([]comment.Comment, int, error) {
	r.logger.Info("Repository: Creating HTTP request for comments")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/comments", r.baseURL), nil)
	if err != nil {
		r.logger.Errorf("Repository: Error creating HTTP request: %v", err)
		return nil, 0, err
	}

	return r.executeRequest(req)
}

func (r *CommentRepository) FetchByPostID(ctx context.Context, postID string) ([]comment.Comment, int, error) {
	r.logger.Infof("Repository: Creating HTTP request for comments of post %s", postID)
	url := fmt.Sprintf("%s/posts/%s/comments", r.baseURL, postID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		r.logger.Errorf("Repository: Error creating HTTP request: %v", err)
		return nil, 0, err
	}

	return r.executeRequest(req)
}

func (r *CommentRepository) executeRequest(req *http.Request) ([]comment.Comment, int, error) {
	r.logger.Info("Repository: Executing HTTP request")
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Errorf("Repository: Error executing HTTP request: %v", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := "repository: Failed to fetch comments: non-OK HTTP status"
		r.logger.Error(errMsg)
		return nil, resp.StatusCode, errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		r.logger.Errorf("Repository: Error reading response body: %v", err)
		return nil, resp.StatusCode, err
	}

	var comments []comment.Comment
	if err := json.Unmarshal(body, &comments); err != nil {
		r.logger.Errorf("Repository: Error unmarshaling comments: %v", err)
		return nil, resp.StatusCode, err
	}

	// Validate each comment
	for i, c := range comments {
		if err := validator.ValidateModel(&c,
			validation.Field(&c.PostID, validation.Required),
			validation.Field(&c.ID, validation.Required),
			validation.Field(&c.Name, validation.Required),
			validation.Field(&c.Email, validation.Required),
			validation.Field(&c.Body, validation.Required),
		); err != nil {
			r.logger.Errorf("Repository: Validation error in comment index %d: %v", i, err)
			return nil, resp.StatusCode, err
		}
	}

	r.logger.Info("Repository: Successfully validated and marshaled comments data")
	return comments, resp.StatusCode, nil
}
