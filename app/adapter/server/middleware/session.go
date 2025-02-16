package middleware

import (
	"github.com/labstack/echo/v4"
)

func (m *MiddlewareFunc) ValidateSession() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sessionCookie, err := c.Cookie("session_id")
			if err != nil {
				return c.JSON(401, "Unauthorized")
			}
			sessionID := sessionCookie.Value
			user, err := m.cacheClient.GetSession(c.Request().Context(), sessionID)
			if err != nil {
				return c.JSON(401, "Unauthorized")
			}
			c.Set("user", user)
			return next(c)
		}
	})
}
