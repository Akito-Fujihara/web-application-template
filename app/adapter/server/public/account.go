package public

import (
	"net/http"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/labstack/echo/v4"
)

func (s *PublicServer) SignUpCreate(c echo.Context) error {
	ctx := c.Request().Context()
	defer c.SetRequest(c.Request().WithContext(ctx))

	var body publicapi.SignUpRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	user := &model.User{
		Name : body.Username,
		Email: body.Email,
	}

	if err := s.accountUsecase.CreateUser(ctx, user, body.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
