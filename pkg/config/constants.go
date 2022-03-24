package config

import "time"

const (
	Development string = "dev"
	Production  string = "prod"
)

var (
	HttpTimeOut           time.Duration = 4 * time.Second
	ServerShutdownTimeOut time.Duration = 10 * time.Second
)
