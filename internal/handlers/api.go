package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	api "github.com/suchithsridhar/suchicodes/web/views/api"
	cmpts "github.com/suchithsridhar/suchicodes/web/views/components"
)

func (h* Handler) handleApiHome(ctx echo.Context) error {
	apiData, err := loadJSONFromFile[api.ApiHomeJSON](api.ApiHomeJsonFile)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling API load.")
	}

	navbarData, err := loadJSONFromFile[cmpts.NavbarJSON](cmpts.NavbarJsonFile)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling API load.")
	}

	footerData, err := loadJSONFromFile[cmpts.FooterJSON](cmpts.FooterJsonFile)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling API load.")
	}

	return renderTemplate(ctx, api.ApiShow(apiData, navbarData, footerData))
}

func (h* Handler) handleApiApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[api.ApiHomeJSON](api.ApiHomeJsonFile)
	return ctx.JSON(status, data)
}
