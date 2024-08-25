package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	cookie "github.com/suchithsridhar/suchicodes/pkg/cookie_manager"
	cmpts "github.com/suchithsridhar/suchicodes/web/views/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderTemplate(c echo.Context, component templ.Component) error {
	ctxWithUser := context.WithValue(
		c.Request().Context(),
		cookie.UserKey,
		c.Get(cookie.UserKey),
	)

	ctx := context.WithValue(
		ctxWithUser,
		cookie.ThemeKey,
		c.Get(cookie.ThemeKey),
	)

	return component.Render(ctx, c.Response())
}

func loadJSONFromFile[T any](filename string) (*T, error) {
	var data T

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func serverJSONAsApi[T any](filename string) (int, interface{}) {
	data, err := loadJSONFromFile[T](filename)

	if err != nil {
		return http.StatusInternalServerError, echo.Map{
			"error": "Unable to access or parse " + filename,
			"value": err.Error(),
		}
	}

	return http.StatusOK, data
}

func getNavbarAndFooter() (*cmpts.NavbarJSON, *cmpts.FooterJSON, error) {
	navbarData, err := loadJSONFromFile[cmpts.NavbarJSON](cmpts.NavbarJsonFile)
	if err != nil {
		return nil, nil, err
	}

	footerData, err := loadJSONFromFile[cmpts.FooterJSON](cmpts.FooterJsonFile)
	if err != nil {
		return nil, nil, err
	}

	return navbarData, footerData, nil
}
