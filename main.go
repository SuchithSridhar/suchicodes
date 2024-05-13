package main

import (
	"fmt"
	"log/slog"

	echo "github.com/labstack/echo/v4"
	configs "github.com/suchithsridhar/suchicodes/configs"
	"github.com/suchithsridhar/suchicodes/internal/handlers"
)

func main() {
	configs.InitConfig()

	app := echo.New()

	handlers.SetupHandlers(app)

	slog.Info(fmt.Sprintf("Starting %s server on port %d", configs.Config.ENVIRONMENT, configs.Config.LISTEN_PORT))
	app.Start(fmt.Sprintf(":%d", configs.Config.LISTEN_PORT))
}
