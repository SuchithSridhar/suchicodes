package handlers

import (
	echo "github.com/labstack/echo/v4"
	components "github.com/suchithsridhar/suchicodes/web/views/components"
)

func handleNavbarApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[components.NavbarJSON](components.NavbarJsonFile)
	return ctx.JSON(status, data)
}

func handleFooterApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[components.FooterJSON](components.FooterJsonFile)
	return ctx.JSON(status, data)
}
