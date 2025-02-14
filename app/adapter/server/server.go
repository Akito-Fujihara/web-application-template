package server

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/middleware"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/privateapi"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql/ormgen"
)

var Set = wire.NewSet(
	NewApplication,
)

type Application struct {
	ApiServer *echo.Echo
	Middleware *middleware.MiddlewareFunc
	DB        *gorm.DB
}

func NewApplication(
	m *middleware.MiddlewareFunc,
	privateServer privateapi.ServerInterface,
	publicServer publicapi.ServerInterface,
	db *gorm.DB,
) *Application {
	e := echo.New()

	e.Use(m.CORS)

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
