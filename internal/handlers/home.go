package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	views_home "github.com/suchithsridhar/suchicodes/web/views/home"
)

const jsonFile = "./web/views/home/index.json"

type IndexJSON struct {
	Title          string `json:title`
	LandingTitle   string `json:"landing_title"`
	LandingAbout   string `json:"landing_about"`
	LandingBtn1    string `json:"landing_btn1"`
	LandingBtn2    string `json:"landing_btn2"`
	LandingBtn1URL string `json:"landing_btn1_url"`
	LandingBtn2URL string `json:"landing_btn2_url"`
}

func handleIndexShow(ctx echo.Context) error {
	return renderTemplate(ctx, views_home.Index())
}

func handleIndexApi(ctx echo.Context) error {
	indexData, err := loadJSONFromFile[IndexJSON](jsonFile)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Unable to access or parse index.json",
			"value": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, indexData)
}
