package depx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	validator "social/depx/validation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	log "github.com/sirupsen/logrus"
)

type DepxComment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func FetchComments(ctx context.Context) ([]DepxComment, int, error) {
	log.Info("Depx: Creating HTTP request for comments")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/comments", nil)
	if err != nil {
		log.Errorf("Depx: Error creating HTTP request: %v", err)
		return nil, 0, err
	}

	return executeCommentRequest(req)
}

func FetchCommentsByPostID(ctx context.Context, postID string) ([]DepxComment, int, error) {
	log.Infof("Depx: Creating HTTP request for comments of post %s", postID)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s/comments", postID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("Depx: Error creating HTTP request: %v", err)
		return nil, 0, err
	}

	return executeCommentRequest(req)
}

func executeCommentRequest(req *http.Request) ([]DepxComment, int, error) {
	log.Info("Depx: Executing HTTP request")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("Depx: Error executing HTTP request: %v", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := "depx: Failed to fetch comments: non-OK HTTP status"
		log.Error(errMsg)
		return nil, resp.StatusCode, errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Depx: Error reading response body: %v", err)
		return nil, resp.StatusCode, err
	}

	var comments []DepxComment
	if err := json.Unmarshal(body, &comments); err != nil {
		log.Errorf("Depx: Error unmarshaling comments: %v", err)
		return nil, resp.StatusCode, err
	}

	// Validate each comment
	for i, comment := range comments {
		if err := validator.ValidateModel(&comment,
			validation.Field(&comment.PostID, validation.Required),
			validation.Field(&comment.ID, validation.Required),
			validation.Field(&comment.Name, validation.Required),
			validation.Field(&comment.Email, validation.Required),
			validation.Field(&comment.Body, validation.Required),
		); err != nil {
			log.Errorf("Depx: Validation error in comment index %d: %v", i, err)
			return nil, resp.StatusCode, err
		}
	}

	log.Info("Depx: Successfully validated and marshaled comments data")
	return comments, resp.StatusCode, nil
}
