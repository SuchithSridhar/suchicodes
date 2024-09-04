package handlers

import (
	echo "github.com/labstack/echo/v4"
	cmpts "github.com/suchithsridhar/suchicodes/web/views/components"
	api "github.com/suchithsridhar/suchicodes/web/views/api"
)

func (h* Handler) handleApiHome(ctx echo.Context) error {
	apiData, err := loadJSONFromFile[api.ApiHomeJSON](api.ApiHomeJsonFile)
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

	return renderTemplate(ctx, api.ApiShow(apiData, navbarData, footerData))
}

func (h* Handler) handleApiApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[api.ApiHomeJSON](api.ApiHomeJsonFile)
	return ctx.JSON(status, data)
}
