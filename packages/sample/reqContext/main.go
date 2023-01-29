package main

import (
    "fmt"
)

type Response struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       interface{}       `json:"body,omitempty"`
}

func Main(r map[string]string) (*Response, error) {
    fmt.Println(r["__ow_method"])
    return &Response{Body: "hello, World"}, nil
}
