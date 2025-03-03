package providers

import (
	"context"
	"social/depx"

	log "github.com/sirupsen/logrus"
)

// Comment represents the shape of a comment in the provider layer
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// GetComments calls the depx layer to fetch all comments
func GetComments(ctx context.Context) ([]Comment, int, error) {
	log.Info("Provider: Calling depx.FetchComments")
	depxComments, status, err := depx.FetchComments(ctx)
	if err != nil {
		log.Errorf("Provider: Error fetching comments: %v", err)
		return nil, status, err
	}

	var comments []Comment
	for _, dc := range depxComments {
		comments = append(comments, Comment{
			PostID: dc.PostID,
			ID:     dc.ID,
			Name:   dc.Name,
			Email:  dc.Email,
			Body:   dc.Body,
		})
	}

	log.Info("Provider: Successfully converted comments to typed response")
	return comments, status, nil
}

// GetCommentsByPostID calls the depx layer to fetch comments for a specific post
func GetCommentsByPostID(ctx context.Context, postID string) ([]Comment, int, error) {
	log.Infof("Provider: Calling depx.FetchCommentsByPostID for post %s", postID)
	depxComments, status, err := depx.FetchCommentsByPostID(ctx, postID)
	if err != nil {
		log.Errorf("Provider: Error fetching comments for post %s: %v", postID, err)
		return nil, status, err
	}

	var comments []Comment
	for _, dc := range depxComments {
		comments = append(comments, Comment{
			PostID: dc.PostID,
			ID:     dc.ID,
			Name:   dc.Name,
			Email:  dc.Email,
			Body:   dc.Body,
		})
	}

	log.Info("Provider: Successfully converted comments to typed response")
	return comments, status, nil
}
