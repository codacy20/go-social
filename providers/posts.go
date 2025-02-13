package providers

import (
	"context"
	"encoding/json"
	"social/depx"

	log "github.com/sirupsen/logrus"
)

// Post represents the shape of a post.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// GetPosts calls the depx layer to fetch posts data,
// converts the raw JSON to a slice of Post structs,
// and marshals it back to JSON.
func GetPosts(ctx context.Context) ([]byte, int, error) {
	log.Info("Provider: Calling depx.FetchPosts")
	data, status, err := depx.FetchPosts(ctx)
	if err != nil {
		log.Errorf("Provider: Error fetching posts: %v", err)
		return nil, status, err
	}

	log.Info("Provider: Unmarshaling posts data")
	var posts []Post
	if err := json.Unmarshal(data, &posts); err != nil {
		log.Errorf("Provider: Error unmarshaling posts: %v", err)
		return nil, status, err
	}

	log.Infof("Provider: Successfully unmarshaled %d posts", len(posts))

	finalData, err := json.Marshal(posts)
	if err != nil {
		log.Errorf("Provider: Error marshaling posts: %v", err)
		return nil, status, err
	}
	log.Info("Provider: Successfully marshaled posts data")
	return finalData, status, nil
}
