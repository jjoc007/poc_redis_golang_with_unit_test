package config

import (
	"time"
)

type (
	Redis struct {
		RedisURL      string
		RedisPassword string
		Timeout       time.Duration
		PoolSize      int
		DialTimeout   time.Duration
	}
)
