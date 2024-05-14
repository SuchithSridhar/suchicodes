package handlers

import (
	echo "github.com/labstack/echo/v4"

	"github.com/suchithsridhar/suchicodes/configs"
)

func SetupHandlers(app *echo.Echo) {

	app.GET("/", handleIndexShow)
	app.GET("/api/pages/index", handleIndexApi)

	// components
	app.GET("/api/components/navbar", handleNavbarApi)
	app.GET("/api/components/footer", handleFooterApi)

	// static files
	app.Static("/assets", configs.Config.ASSETS_DIR)
	app.Static("/css", configs.Config.STYLES_DIR)
	app.Static("/js", configs.Config.JAVASCRIPT_DIR)
}
