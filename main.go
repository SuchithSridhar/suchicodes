package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/suchithsridhar/suchicodes/app"
	"github.com/suchithsridhar/suchicodes/internal/config"
	"github.com/suchithsridhar/suchicodes/internal/database"
	"github.com/suchithsridhar/suchicodes/internal/handlers"
	"github.com/suchithsridhar/suchicodes/internal/logger"
	"github.com/suchithsridhar/suchicodes/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.InitConfig()
	logger, logFile := logger.InitLogger(cfg)
	defer logFile.Close()

	db, err := gorm.Open(sqlite.Open(cfg.DATABASE_URI), &gorm.Config{})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to open SQLite database: %v\n", err))
		os.Exit(1)
	} else {
		logger.Info(fmt.Sprintf("Using SQLite database at %v", cfg.DATABASE_URI))
	}

	if err := database.InitializeDatabase(db, logger); err != nil {
		logger.Error(fmt.Sprintf("Database initialization failed: %v", err))
		os.Exit(1)
	}

	server := echo.New()

	app := &app.App{
		Server: server,
		Db:     db,
		Config: cfg,
		Logger: logger,
	}

	server.Use(middleware.RealIPMiddleware())

	handlers.SetupHandlers(app)

	logger.Info(fmt.Sprintf("Starting %s server on port %d", cfg.ENVIRONMENT, cfg.LISTEN_PORT))
	server.Start(fmt.Sprintf(":%d", cfg.LISTEN_PORT))
}
