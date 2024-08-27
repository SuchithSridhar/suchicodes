package config

import (
	"log/slog"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GlobalConfig struct {
	ENVIRONMENT    string
	SECRET_KEY     []byte
	DATABASE_URI   string
	LOG_FILE       string
	LOG_LEVEL      string
	LOG_ADD_SOURCE bool
	ROOT_DIRECTORY string
	LISTEN_PORT    int
	ASSETS_DIR     string
	STYLES_DIR     string
	JAVASCRIPT_DIR string
}

const (
	development = "development"
	production  = "production"
	listenPortMin = 20
	listenPortMax = 60000
)

var (
	development_secret_key = "secret"
	validLogLevels         = map[string]struct{}{
		"DEBUG": {},
		"INFO":  {},
		"WARN":  {},
		"ERROR": {},
	}
)

func InitConfig() *GlobalConfig {
	config := &GlobalConfig{
		ENVIRONMENT:    development,
		SECRET_KEY:     []byte(development_secret_key),
		DATABASE_URI:   "",
		LOG_FILE:       "stdout",
		LOG_LEVEL:      "DEBUG",
		LOG_ADD_SOURCE: true,
		ROOT_DIRECTORY: "",
		LISTEN_PORT:    3000,
		ASSETS_DIR:     "./web/static/assets",
		STYLES_DIR:     "./web/static/css",
		JAVASCRIPT_DIR: "./web/static/js",
	}

	var value string
	var ok bool

	if value, ok = os.LookupEnv("ENVIRONMENT"); ok {
		config.ENVIRONMENT = value
	} else {
		slog.Info(fmt.Sprintf("ENVIRONMENT variable not set, defaulting to %v", config.ENVIRONMENT))
	}

	if value, ok = os.LookupEnv("SECRET_KEY"); ok {
		config.SECRET_KEY = []byte(value)
	} else if config.ENVIRONMENT != development {
		slog.Error("Secret key not set in non-development environment.")
		os.Exit(1)
	}

	if value, ok = os.LookupEnv("DATABASE_URI"); ok {
		config.DATABASE_URI = value
	} else {
		slog.Error("DATABASE_URI not set, exiting.")
		os.Exit(1)
	}

	if value, ok = os.LookupEnv("LOG_FILE"); ok {
		config.LOG_FILE = value
	} else {
		slog.Warn(fmt.Sprintf("LOG_FILE not set, defaulting to %v.", config.LOG_FILE))
		os.Exit(1)
	}

	if value, ok = os.LookupEnv("LOG_LEVEL"); ok {
		value = strings.ToUpper(value)
		if _, exists := validLogLevels[value]; !exists {
			slog.Error("Invalid LOG_LEVEL set!")
			os.Exit(1)
		}
		config.LOG_LEVEL = value
	} else {
		slog.Warn(fmt.Sprintf("LOG_LEVEL not set, defaulting to %v.", config.LOG_LEVEL))
	}

	if value, ok = os.LookupEnv("LOG_ADD_SOURCE"); ok {
		switch value {
		case "true":
			config.LOG_ADD_SOURCE = true
		case "false":
			config.LOG_ADD_SOURCE = false
		default:
			slog.Error(fmt.Sprintf("Invalid value for LOG_ADD_SOURCE [true/false]: %v.", value))
			os.Exit(1)
		}
	} else {
		slog.Info(fmt.Sprintf("LOG_ADD_SOURCE not set, defaulting to %v.", config.LOG_ADD_SOURCE))
	}

	if value, ok = os.LookupEnv("ROOT_DIRECTORY"); ok {
		config.ROOT_DIRECTORY = value
	} else if config.ENVIRONMENT != development {
		slog.Error("Root directory not set in non-development environment.")
		os.Exit(1)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			slog.Error("Unable to set root directory.")
			os.Exit(1)
		}
		config.ROOT_DIRECTORY = cwd
	}

	if value, ok := os.LookupEnv("LISTEN_PORT"); ok {
		intvalue, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil || intvalue < listenPortMin || intvalue > listenPortMax {
			slog.Error(fmt.Sprintf("Invalid value set for LISTEN_PORT: %v.", intvalue))
			os.Exit(1)
		}

		config.LISTEN_PORT = intvalue
	} else {
		slog.Info(fmt.Sprintf("LISTEN_PORT not set, defaulting to %v.", config.LISTEN_PORT))
	}

	return config
}
