package app

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/internal/config"
	"log/slog"
)

type App struct {
	Server *echo.Echo
	Db     *sql.DB
	Config *config.GlobalConfig
	Logger *slog.Logger
}
