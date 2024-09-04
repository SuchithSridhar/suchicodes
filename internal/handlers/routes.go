package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/app"
)

type Handler struct {
	App *app.App
}

func SetupHandlers(app *app.App) {
	handler := &Handler{App: app}

	server := app.Server
	config := app.Config

	// home pages
	server.GET("/", handler.handleIndexShow)
	routeWithoutSlash(server, "/about", "get", handler.handleAboutShow)
	routeWithoutSlash(server, "/contact", "get", handler.handleContactShow)

	// post requests pages
	routeWithoutSlash(server, "/contact", "post", handler.handleContactPost)

	// api endpoints
	routeWithSlash(server, "/api/", "get", handler.handleApiHome)
	routeWithoutSlash(server, "/api/pages/index", "get", handler.handleIndexApi)
	routeWithoutSlash(server, "/api/pages/about", "get", handler.handleAboutApi)
	routeWithoutSlash(server, "/api/pages/contact", "get", handler.handleContactApi)
	routeWithoutSlash(server, "/api/pages/api", "get", handler.handleApiApi)

	// components
	routeWithoutSlash(server, "/api/components/navbar", "get", handler.handleNavbarApi)
	routeWithoutSlash(server, "/api/components/footer", "get", handler.handleFooterApi)

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

	if method == "get" {
		e.GET(pathWithSlash, handler)
		e.GET(pathWithoutSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithSlash)
		})
	} else if method == "post" {
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

	if method == "get" {
		e.GET(pathWithoutSlash, handler)
		e.GET(pathWithSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithoutSlash)
		})
	} else if method == "post" {
		e.POST(pathWithoutSlash, handler)
		e.POST(pathWithSlash, func(c echo.Context) error {
			return c.Redirect(http.StatusMovedPermanently, pathWithoutSlash)
		})
	}
}
