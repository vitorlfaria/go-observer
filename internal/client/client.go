package client

import (
	"net/http"
	"time"
)

func NewClient() http.Client {
	return http.Client{
		Timeout: 5 * time.Second,
	}
}
