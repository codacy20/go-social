package depx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	validator "social/depx/validation"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	log "github.com/sirupsen/logrus"
)

// Create a custom HTTP client with a timeout.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type DepxPost struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func FetchPosts(ctx context.Context) ([]byte, int, error) {
	log.Info("Depx: Creating HTTP request for posts")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		log.Errorf("Depx: Error creating HTTP request: %v", err)
		return nil, 0, err
	}

	log.Info("Depx: Executing HTTP request")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("Depx: Error executing HTTP request: %v", err)
		return nil, 0, err
	}
	defer resp.Body.Close()
	log.Infof("Depx: Received response with status code %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		errMsg := "depx: Failed to fetch posts: non-OK HTTP status"
		log.Error(errMsg)
		return nil, resp.StatusCode, errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Depx: Error reading response body: %v", err)
		return nil, resp.StatusCode, err
	}
	log.Info("Depx: Successfully read response body")

	// Unmarshal the raw JSON into a slice of DepxPost.
	var posts []DepxPost
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Errorf("Depx: Error unmarshaling posts: %v", err)
		return nil, resp.StatusCode, err
	}

	// Validate each post.
	for i, post := range posts {
		if err := validator.ValidateModel(&post,
			validation.Field(&post.UserID, validation.Required),
			validation.Field(&post.ID, validation.Required),
			validation.Field(&post.Title, validation.Required),
			validation.Field(&post.Body, validation.Required),
		); err != nil {
			log.Errorf("Depx: Validation error in post index %d: %v", i, err)
			return nil, resp.StatusCode, err
		}
	}

	// Marshal the validated posts back to JSON.
	validatedData, err := json.Marshal(posts)
	if err != nil {
		log.Errorf("Depx: Error marshaling validated posts: %v", err)
		return nil, resp.StatusCode, err
	}

	log.Info("Depx: Successfully validated and marshaled posts data")
	return validatedData, resp.StatusCode, nil
}
