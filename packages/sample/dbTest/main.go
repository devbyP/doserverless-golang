package main

import (
    "os"
    "context"
)

type Response struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       interface{}       `json:"body,omitempty"`
}

func Main() (*Response, error) {
	url := os.Getenv("DATABASE_URL")
    ctx := context.Background()
    _, err := connectDatabase(ctx, url)
    if err != nil {
        return nil, err
    }
	return &Response{}, nil
}
