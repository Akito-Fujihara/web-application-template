package server

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/privateapi"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/ormgen"
)

var Set = wire.NewSet(
	NewApplication,
)

type Application struct {
	ApiServer *echo.Echo
	DB        *gorm.DB
}

func NewApplication(
	privateServer privateapi.ServerInterface,
	publicServer publicapi.ServerInterface,
	db *gorm.DB,
) *Application {
	e := echo.New()

	privateapi.RegisterHandlers(e, privateServer)
	publicapi.RegisterHandlers(e, publicServer)

	return &Application{
		ApiServer: e,
		DB:        db,
	}
}

func (a *Application) Start() error {
	ormgen.SetDefault(a.DB)
	return a.ApiServer.Start(":8080")
}
