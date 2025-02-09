package server

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/privateapi"
	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
)

var Set = wire.NewSet(
	NewApplication,
	NewPrivateServer,
	NewPublicAServer,
)

type Application struct {
	ApiServer *echo.Echo
}

func NewApplication(privateServer *PrivateServer, publicServer *PublicServer) *Application {
	e := echo.New()

	privateapi.RegisterHandlers(e, privateServer)
	publicapi.RegisterHandlers(e, publicServer)

	return &Application{
		ApiServer: e,
	}
}

func (a *Application) Start() error {
	return a.ApiServer.Start(":8080")
}

type PrivateServer struct{}

func NewPrivateServer() *PrivateServer {
	return &PrivateServer{}
}

var _ privateapi.ServerInterface = (*PrivateServer)(nil)

type PublicServer struct{}

func NewPublicAServer() *PublicServer {
	return &PublicServer{}
}

var _ publicapi.ServerInterface = (*PublicServer)(nil)
