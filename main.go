package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/configs"
	"github.com/suchithsridhar/suchicodes/internal/handlers"
)

func main() {
	configs.InitConfig()

	app := echo.New()

	handlers.SetupHandlers(app)

	app.Start(":3000")
}
