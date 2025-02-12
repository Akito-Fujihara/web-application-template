package private

import (
	"net/http"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/privateapi"
	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/labstack/echo/v4"
)

func (s *PrivateServer) CreateTodo(c echo.Context) error {
	ctx := c.Request().Context()
	defer c.SetRequest(c.Request().WithContext(ctx))

	var body privateapi.CreateTodoRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var todo model.Todo
	todo.UserID = 1 // 一旦 middleware が入るまでの暫定
	todo.Title = body.Title
	todo.Description = &body.Description

	if err := s.todoUsercase.CreateTodo(ctx, &todo); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, nil)
}
