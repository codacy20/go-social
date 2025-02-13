package providers

import (
	"context"
	"encoding/json"
	"social/depx"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts(ctx context.Context) ([]byte, int, error) {
	data, status, err := depx.FetchPosts(ctx)
	if err != nil {
		return nil, status, err
	}

	var posts []Post
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, status, err
	}

	finalData, err := json.Marshal(posts)
	if err != nil {
		return nil, status, err
	}

	return finalData, status, nil
}
