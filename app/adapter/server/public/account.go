package public

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Akito-Fujihara/web-application-template/app/adapter/server/oapi/publicapi"
	"github.com/Akito-Fujihara/web-application-template/app/domain/model"
	"github.com/labstack/echo/v4"
)

func (s *PublicServer) SignUp(c echo.Context) error {
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

	return c.JSON(http.StatusCreated, nil)
}

func (s *PublicServer) Login(c echo.Context) error {
	ctx := c.Request().Context()
	defer c.SetRequest(c.Request().WithContext(ctx))

	var body publicapi.LoginRequest
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	sessionID, err := s.accountUsecase.AuthenticateUser(ctx, body.Email, body.Password)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to authenticate user: %v", err))
		return c.JSON(http.StatusUnauthorized, nil)
	}

	// セッションIDをCookieにセットし、1日後に削除されるように設定
	c.SetCookie(&http.Cookie{
		Name: "session_id",
		Value: sessionID,
		Path: "/",
		Domain: "localhost", // 一旦ベタ書きだが、環境変数などで設定する
		Expires: time.Now().Add(24 * time.Hour),
		// Secure: true, // 本番環境で有効にする
		HttpOnly: true,
	})

	// フロントエンドでログイン状態を保持するためのCookieをセット
	c.SetCookie(&http.Cookie{
		Name: "signed_in",
		Value: "true",
		Path: "/",
		Domain: "localhost", // 一旦ベタ書きだが、環境変数などで設定する
		Expires: time.Now().Add(24 * time.Hour),
		// Secure: true, // 本番環境で有効にする
		HttpOnly: false,
	})

	return c.JSON(http.StatusOK, nil)
}
