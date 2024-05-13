package handlers

import (
	echo "github.com/labstack/echo/v4"
)

func SetupHandlers(app *echo.Echo) {

	app.GET("/", handleIndexShow)
	app.GET("/api/pages/index", handleIndexApi)

	// components
	app.GET("/api/components/navbar", handleNavbarApi)
	app.GET("/api/components/footer", handleFooterApi)
}
