package types

import "time"

type Config struct {
	URLs []string `json:"urls"`
}

type Result struct {
	URL      string
	Status   string
	Duration time.Duration
	Error    error
}
