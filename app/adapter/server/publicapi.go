package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *PublicServer) SignUpCreate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
