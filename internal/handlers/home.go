package handlers

import (
	echo "github.com/labstack/echo/v4"
	home "github.com/suchithsridhar/suchicodes/web/views/home"
)

func handleIndexShow(ctx echo.Context) error {
	indexData, err := loadJSONFromFile[home.IndexJSON](home.IndexJsonFile)

	if err != nil {
		// handle error
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		// handle error
	}

	return renderTemplate(ctx, home.IndexShow(indexData, navbarData, footerData))
}

func handleAboutShow(ctx echo.Context) error {
	aboutData, err := loadJSONFromFile[home.AboutJSON](home.AboutJsonFile)

	if err != nil {
		// handle error
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		// handle error
	}

	return renderTemplate(ctx, home.AboutShow(aboutData, navbarData, footerData))
}

func handleContactShow(ctx echo.Context) error {
	contactData, err := loadJSONFromFile[home.ContactJSON](home.ContactJsonFile)

	if err != nil {
		// handle error
	}

	navbarData, footerData, err := getNavbarAndFooter()

	if err != nil {
		// handle error
	}

	return renderTemplate(ctx, home.ContactShow(contactData, navbarData, footerData))
}

func handleIndexApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.IndexJSON](home.IndexJsonFile)
	return ctx.JSON(status, data)
}

func handleAboutApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.AboutJSON](home.AboutJsonFile)
	return ctx.JSON(status, data)
}

func handleContactApi(ctx echo.Context) error {
	status, data := serverJSONAsApi[home.ContactJSON](home.ContactJsonFile)
	return ctx.JSON(status, data)
}
