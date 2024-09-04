package app

import (
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/internal/config"
	"log/slog"
)

type App struct {
	Server *echo.Echo
	Db     *gorm.DB
	Config *config.GlobalConfig
	Logger *slog.Logger
}
