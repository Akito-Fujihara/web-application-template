package cache

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

var Set = wire.NewSet(
	NewCacheConn,
)

type Config struct {
	Host string
	Port int
	Password string
}

func NewCacheConn(config *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB: 0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}
	return client, nil
}
