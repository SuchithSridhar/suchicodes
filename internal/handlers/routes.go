package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/app"
)

func SetupHandlers(app *app.App) {

	server := app.Server
	config := app.Config

	// home pages
	server.GET("/", handleIndexShow)
	routeWithoutSlash(server, "/about", "get", handleAboutShow)
	routeWithoutSlash(server, "/contact", "get", handleContactShow)

	// api endpoints
	routeWithSlash(server, "/api/", "get", handleApiHome)
	routeWithoutSlash(server, "/api/pages/index", "get", handleIndexApi)
	routeWithoutSlash(server, "/api/pages/about", "get", handleAboutApi)
	routeWithoutSlash(server, "/api/pages/contact", "get", handleContactApi)
	routeWithoutSlash(server, "/api/pages/api", "get", handleApiApi)

	// components
	routeWithoutSlash(server, "/api/components/navbar", "get", handleNavbarApi)
	routeWithoutSlash(server, "/api/components/footer", "get", handleFooterApi)

	// static files
	server.Static("/assets", config.ASSETS_DIR)
	server.Static("/css", config.STYLES_DIR)
	server.Static("/js", config.JAVASCRIPT_DIR)
}

func routeWithSlash(e *echo.Echo, path string, method string, handler echo.HandlerFunc) {
	var pathWithSlash string
	var pathWithoutSlash string

	if !strings.HasSuffix(path, "/") {
		pathWithSlash = path + "/"
		pathWithoutSlash = path
	} else {
		pathWithSlash = path
		pathWithoutSlash = strings.TrimSuffix(path, "/")
	}
	
	if (method == "get") {
		e.GET(pathWithSlash, handler)
		e.GET(pathWithoutSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithSlash)
		})
	} else if (method == "post") {
		e.POST(pathWithSlash, handler)
		e.POST(pathWithoutSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithSlash)
		})
	}
}

func routeWithoutSlash(e *echo.Echo, path string, method string, handler echo.HandlerFunc) {
	var pathWithSlash string
	var pathWithoutSlash string

	if !strings.HasSuffix(path, "/") {
		pathWithSlash = path + "/"
		pathWithoutSlash = path
	} else {
		pathWithSlash = path
		pathWithoutSlash = strings.TrimSuffix(path, "/")
	}

	if (method == "get") {
		e.GET(pathWithoutSlash, handler)
		e.GET(pathWithSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithoutSlash)
		})
	} else if (method == "post") {
		e.POST(pathWithoutSlash, handler)
		e.POST(pathWithSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithoutSlash)
		})
	}
}
