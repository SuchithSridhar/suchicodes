package handlers

import (
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"

	"github.com/suchithsridhar/suchicodes/configs"
)

func SetupHandlers(app *echo.Echo) {


	// home pages
	app.GET("/", handleIndexShow)
	routeWithoutSlash(app, "/about", handleAboutShow)
	routeWithoutSlash(app, "/contact", handleContactShow)

	// api endpoints
	routeWithSlash(app, "/api/", handleApiHome)
	routeWithoutSlash(app, "/api/pages/index", handleIndexApi)
	routeWithoutSlash(app, "/api/pages/about", handleAboutApi)
	routeWithoutSlash(app, "/api/pages/contact", handleContactApi)
	routeWithoutSlash(app, "/api/pages/api", handleApiApi)

	// components
	routeWithoutSlash(app, "/api/components/navbar", handleNavbarApi)
	routeWithoutSlash(app, "/api/components/footer", handleFooterApi)

	// static files
	app.Static("/assets", configs.Config.ASSETS_DIR)
	app.Static("/css", configs.Config.STYLES_DIR)
	app.Static("/js", configs.Config.JAVASCRIPT_DIR)
}

func routeWithSlash(e *echo.Echo, path string, handler echo.HandlerFunc) {
	var pathWithSlash string
	var pathWithoutSlash string

	if !strings.HasSuffix(path, "/") {
		pathWithSlash = path + "/"
		pathWithoutSlash = path
	} else {
		pathWithSlash = path
		pathWithoutSlash = strings.TrimSuffix(path, "/")
	}
	
	e.GET(pathWithSlash, handler)
	e.GET(pathWithoutSlash, func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, pathWithSlash)
	})
}

func routeWithoutSlash(e *echo.Echo, path string, handler echo.HandlerFunc) {
	var pathWithSlash string
	var pathWithoutSlash string

	if !strings.HasSuffix(path, "/") {
		pathWithSlash = path + "/"
		pathWithoutSlash = path
	} else {
		pathWithSlash = path
		pathWithoutSlash = strings.TrimSuffix(path, "/")
	}

	e.GET(pathWithoutSlash, handler)
	e.GET(pathWithSlash, func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, pathWithoutSlash)
	})
}
