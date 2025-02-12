// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/private"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/public"
	"github.com/Akito-Fujihara/web-application-template/app/config/env"
	"github.com/Akito-Fujihara/web-application-template/app/infra/cache"
	"github.com/Akito-Fujihara/web-application-template/app/infra/cache/cacheclient"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/repository"
	"github.com/Akito-Fujihara/web-application-template/app/usecase"
)

// Injectors from wire.go:

func InitializeServer() (*server.Application, func(), error) {
	iTodoRepository := repository.NewTodoRepository()
	iTodoUsecase := usecase.NewTodoUsecase(iTodoRepository)
	serverInterface := private.NewPrivateServer(iTodoUsecase)
	config, err := env.CacheConfig()
	if err != nil {
		return nil, nil, err
	}
	client, err := cache.NewCacheConn(config)
	if err != nil {
		return nil, nil, err
	}
	iCacheClient := cacheclient.NewCacheClient(client)
	iUserRepository := repository.NewUserRepository()
	iAccountUsecase := usecase.NewAccountUsecase(iCacheClient, iUserRepository)
	publicapiServerInterface := public.NewPublicServer(iAccountUsecase)
	mysqlConfig, err := env.MysqlConfig()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := mysql.NewMysqlConn(mysqlConfig)
	if err != nil {
		return nil, nil, err
	}
	application := server.NewApplication(serverInterface, publicapiServerInterface, db)
	return application, func() {
		cleanup()
	}, nil
}
