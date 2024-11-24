package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	notes "github.com/suchithsridhar/suchicodes/web/views/notes"
)

func (h *Handler) handleNotesIndexShow(ctx echo.Context) error {
	indexData, err := loadJSONFromFile[notes.IndexJSON](notes.IndexJsonFile)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error in handling page load.")
	}

	return renderTemplate(ctx, notes.IndexShow(indexData, navbarData, footerData))
}

func (h *Handler) handleNotesIndexApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[notes.IndexJSON](notes.IndexJsonFile)
	return ctx.JSON(status, data)
}
