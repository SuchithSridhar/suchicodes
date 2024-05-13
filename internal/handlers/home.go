package handlers

import (
	echo "github.com/labstack/echo/v4"
	views_home "github.com/suchithsridhar/suchicodes/web/views/home"
)

func handleIndexShow(ctx echo.Context) error {
	return renderTemplate(ctx, views_home.Index())
}
