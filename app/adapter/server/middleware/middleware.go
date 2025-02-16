package middleware

import (
	"github.com/Akito-Fujihara/web-application-template/app/domain/cacheclient"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var Set = wire.NewSet(
	NewMiddlewareFunc,
)

type IMiddlewareFunc interface {
	CORS() echo.MiddlewareFunc
	ValidateSession() echo.MiddlewareFunc
}

type MiddlewareFunc struct {
	corsConfig *CORSConfig
	cacheClient cacheclient.ICacheClient
}

func NewMiddlewareFunc(
	corsConfig *CORSConfig,
	cacheClient cacheclient.ICacheClient,
) IMiddlewareFunc {
	return &MiddlewareFunc{
		corsConfig: corsConfig,
		cacheClient: cacheClient,
	}
}
