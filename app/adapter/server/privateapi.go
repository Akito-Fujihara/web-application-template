package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *PrivateServer) CreateTodoCreate(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
