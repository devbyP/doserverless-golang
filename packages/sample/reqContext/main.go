package main

type Http struct {
	Headers map[string]string `json:"headers,omitempty"`
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
}

type Request struct {
	HTTP Http   `json:"http"`
	Name string `json:"name"`
}

type Response struct {
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       Request           `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	return &Response{
		Body: in,
	}, nil
}
