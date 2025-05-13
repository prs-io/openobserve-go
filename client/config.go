package client

import "time"

type Config struct {
	BaseURL  string
	Username string
	Password string
	Timeout  time.Duration
	Headers  map[string]string
	Retries  int
	Debug    bool
}
