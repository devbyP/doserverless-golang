package main

import (
    "os"
	"fmt"
)

type Request struct {
	Header map[string]string `json:"__ow_headers,omitempty"`
	Method string            `json:"__ow_method,omitempty"`
	Name   string            `json:"name"`
}

type Body struct {
	Name string `json:"name"`
}

type Response struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       Body              `json:"body"`
}

func Main(r Request) (*Response, error) {
    text := os.Getenv("TEXT")
    fmt.Println("test text:", text)
	_, ok := r.Header["Content-Type"]
	if !ok {
		fmt.Println("no content type found in req headers")
	}
    if r.Method == "" {
        fmt.Println("method is empty")
    }
	return &Response{
		Headers:     map[string]string{"Content-Type": "application/json"},
        StatusCode: 200,
        Body: Body{Name: r.Name},
	}, nil
}
