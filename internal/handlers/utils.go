package handlers

import (
	"context"

	cookie "github.com/suchithsridhar/suchicodes/pkg/cookie_manager"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderTemplate(c echo.Context, component templ.Component) error {
	ctxWithUser := context.WithValue(
		c.Request().Context(),
		cookie.UserKey,
		c.Get(cookie.UserKey),
	)

	ctx := context.WithValue(
		ctxWithUser,
		cookie.ThemeKey,
		c.Get(cookie.ThemeKey),
	)

	return component.Render(ctx, c.Response())
}
