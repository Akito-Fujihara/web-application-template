//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

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


func InitializeServer() (*server.Application, func(), error) {
	wire.Build(
		env.Set,
		mysql.Set,
		cache.Set,
		repository.Set,
		cacheclient.Set,
		usecase.Set,
		private.Set,
		public.Set,
		server.Set,
	)
	return nil, nil, nil
}
