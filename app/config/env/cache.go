package env

import (
	"fmt"
	"os"

	"github.com/Akito-Fujihara/web-application-template/app/infra/cache"
)

func CacheConfig() (*cache.Config, error) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return nil, fmt.Errorf("REDIS_HOST is not set")
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		return nil, fmt.Errorf("REDIS_PORT is not set")
	}
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("REDIS_PASSWORD is not set")
	}
	return &cache.Config{
		Host:     host,
		Port:     6379,
		Password: password,
	}, nil
}
