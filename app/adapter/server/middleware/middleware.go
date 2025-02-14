package middleware

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var Set = wire.NewSet(
	NewMiddlewareFunc,
	NewCORS,
)

type MiddlewareFunc struct {
	CORS echo.MiddlewareFunc
}

func NewMiddlewareFunc(cors echo.MiddlewareFunc) *MiddlewareFunc {
	return &MiddlewareFunc{
		CORS: cors,
	}
}
