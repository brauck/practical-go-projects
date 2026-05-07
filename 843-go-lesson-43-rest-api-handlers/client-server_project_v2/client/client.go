package client

import "net/http"

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

func New(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}
