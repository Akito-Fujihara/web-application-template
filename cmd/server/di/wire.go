//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server"
)


func InitializeServer() (*server.Application , error) {
	wire.Build(
		server.Set,
	)
	return &server.Application{}, nil
}
