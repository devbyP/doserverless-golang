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
    headers := make(map[string]string)
    headers["Content-type"] = "application/json"
    return &Response{
        Headers: headers,
        Body: "hello, World",
    }, nil
}
