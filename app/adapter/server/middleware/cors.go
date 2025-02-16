package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	AllowOrigins     []string
}

func (m *MiddlewareFunc) CORS() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     m.corsConfig.AllowOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
}
