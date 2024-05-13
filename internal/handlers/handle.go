package handlers

import (
	echo "github.com/labstack/echo/v4"
)

func SetupHandlers(app *echo.Echo) {
	app.GET("/", handleIndexShow)
}
