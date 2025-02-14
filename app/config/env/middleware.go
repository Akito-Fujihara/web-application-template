package env

import (
	"fmt"
	"os"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/middleware"
)

func CORSConfig() (*middleware.CORSConfig, error) { 
	allowOrignURL := os.Getenv("CORS_ALLOW_ORIGIN_URL")
	if allowOrignURL == "" {
		return nil, fmt.Errorf("CORS_ALLOW_ORIGIN_URL is not set")
	}
	return &middleware.CORSConfig{
		AllowOrigins: []string{allowOrignURL},
	}, nil
}
