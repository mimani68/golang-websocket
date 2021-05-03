package config

import "os"

var (
	REDIS_URL      = os.Getenv("REDIS_URL")
	REDIS_PORT     = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
)
