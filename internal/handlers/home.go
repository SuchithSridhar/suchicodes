package handlers

import (
	echo "github.com/labstack/echo/v4"
	home "github.com/suchithsridhar/suchicodes/web/views/home"
	cmpts "github.com/suchithsridhar/suchicodes/web/components"
)

func handleIndexShow(ctx echo.Context) error {
	indexData, err := loadJSONFromFile[home.IndexJSON](home.IndexJsonFile)
	if err != nil {
		// handle error
	}

	navbarData, err := loadJSONFromFile[cmpts.NavbarJSON](cmpts.NavbarJsonFile)
	if err != nil {
		// handle error
	}

	footerData, err := loadJSONFromFile[cmpts.FooterJSON](cmpts.FooterJsonFile)
	if err != nil {
		// handle error
	}

	return renderTemplate(ctx, home.IndexShow(indexData, navbarData, footerData))
}

func handleIndexApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.IndexJSON](home.IndexJsonFile)
	return ctx.JSON(status, data)
}

