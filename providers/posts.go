package providers

import (
	"context"
	"social/depx"

	log "github.com/sirupsen/logrus"
)

// Post represents the shape of a post in the provider layer.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// GetPosts calls the depx layer to fetch posts data and returns a typed slice of posts.
func GetPosts(ctx context.Context) ([]Post, int, error) {
	log.Info("Provider: Calling depx.FetchPosts")
	// Assume depx.FetchPosts now returns a slice of depx.DepxPost.
	depxPosts, status, err := depx.FetchPosts(ctx)
	if err != nil {
		log.Errorf("Provider: Error fetching posts: %v", err)
		return nil, status, err
	}
	log.Infof("Provider: Successfully fetched %d posts from depx", len(depxPosts))

	// Convert depx.DepxPost to provider.Post.
	var posts []Post
	for _, dp := range depxPosts {
		posts = append(posts, Post{
			UserID: dp.UserID,
			ID:     dp.ID,
			Title:  dp.Title,
			Body:   dp.Body,
		})
	}

	log.Infof("Provider: Successfully converted posts to typed response")
	return posts, status, nil
}
