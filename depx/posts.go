package depx

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Create a custom HTTP client with a timeout.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// FetchPosts fetches JSON posts from the external API using the provided context.
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
		errMsg := "depx: failed to fetch posts: non-OK HTTP status"
		log.Error(errMsg)
		return nil, resp.StatusCode, errors.New(errMsg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Depx: Error reading response body: %v", err)
		return nil, resp.StatusCode, err
	}
	log.Info("Depx: Successfully read response body")
	return body, resp.StatusCode, nil
}
