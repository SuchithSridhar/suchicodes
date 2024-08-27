package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/suchithsridhar/suchicodes/internal/config"
)

const stdout = "stdout"

var logLevelStringToValue = map[string]slog.Level{
	"DEBUG": slog.LevelDebug,
	"INFO":  slog.LevelInfo,
	"WARN":  slog.LevelWarn,
	"ERROR": slog.LevelError,
}

func InitLogger(config *config.GlobalConfig) (*slog.Logger, *os.File) {

	var logger *slog.Logger
	var logFile *os.File

	level, ok := logLevelStringToValue[config.LOG_LEVEL]
	if !ok {
		slog.Error("Invalid LOG_LEVEL during logger initialization.")
		os.Exit(1)
	}

	opts := slog.HandlerOptions{
		AddSource: config.LOG_ADD_SOURCE,
		Level:     level,
	}

	if config.LOG_FILE == stdout {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
	} else {
		slog.Info(fmt.Sprintf("LOG_FILE set, logging to file: %v", config.LOG_FILE))
		logFile, err := os.OpenFile(config.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Failed to open log file: %v\n", err)
			os.Exit(1)
		}

		logger = slog.New(slog.NewTextHandler(logFile, &opts))
	}

	return logger, logFile
}
