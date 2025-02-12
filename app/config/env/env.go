package env

import (
	"fmt"
	"os"

	"github.com/Akito-Fujihara/web-application-template/app/infra/cache"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	MysqlConfig,
	CacheConfig,
)

func MysqlConfig() (*mysql.Config, error) {
	username := os.Getenv("MYSQL_USER")
	if username == "" {
		return nil, fmt.Errorf("MYSQL_USER is not set")
	}
	readonlyUsername := os.Getenv("MYSQL_READONLY_USER")
	if readonlyUsername == "" {
		readonlyUsername = username
	}
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("MYSQL_PASSWORD is not set")
	}
	readonlyPassword := os.Getenv("MYSQL_READONLY_PASSWORD")
	if readonlyPassword == "" {
		readonlyPassword = password
	}
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		return nil, fmt.Errorf("MYSQL_HOST is not set")
	}
	readonlyHost := os.Getenv("MYSQL_READONLY_HOST")
	if readonlyHost == "" {
		readonlyHost = host
	}
	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		return nil, fmt.Errorf("MYSQL_PORT is not set")
	}
	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		return nil, fmt.Errorf("MYSQL_DATABASE is not set")
	}
	return &mysql.Config{
		Database:          database,
		Host:              host,
		ReadonlyHost:      readonlyHost,
		Port:              3306,
		Username:          username,
		ReadonlyUsername:  readonlyUsername,
		Password:          password,
		ReadonlyPassword:  readonlyPassword,
	}, nil
}

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
