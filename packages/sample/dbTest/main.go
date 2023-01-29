package main

import (
	"context"
	"net/http"
	"os"
)

type Response struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       []name            `json:"body"`
}

func Main() (*Response, error) {
	url := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	db, err := connectDatabase(ctx, url)
	if err != nil {
		return &Response{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	names, err := getNames(ctx, db)
	if err != nil {
		return &Response{
			StatusCode: http.StatusBadRequest,
		}, err
	}
	return &Response{
		StatusCode: http.StatusOK,
		Body:       names,
	}, nil
}
