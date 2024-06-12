package configs

import (
	"log/slog"
	"os"
	"strconv"
)

type GlobalConfig struct {
	ENVIRONMENT    string
	SECRET_KEY     []byte
	DATABASE_URI   string
	ROOT_DIRECTORY string
	LISTEN_PORT    int
	ASSETS_DIR     string
	STYLES_DIR     string
	JAVASCRIPT_DIR string
}

const development = "development"
const production = "production"

var secret_key = "secret"

var Config = GlobalConfig{
	ENVIRONMENT:    development,
	SECRET_KEY:     nil,
	DATABASE_URI:   "",
	ROOT_DIRECTORY: "",
	LISTEN_PORT:    3000,
	ASSETS_DIR:     "./web/static/assets",
	STYLES_DIR:     "./web/static/css",
	JAVASCRIPT_DIR: "./web/static/js",
}

func InitConfig() {
	var value string
	var ok bool

	if value, ok = os.LookupEnv("ENVIRONMENT"); ok {
		Config.ENVIRONMENT = value
	}

	if value, ok = os.LookupEnv("DATABASE_URI"); ok {
		Config.DATABASE_URI = value
	} else if Config.ENVIRONMENT != development {
		slog.Error("Database URI not set in non-development environment.")
		os.Exit(1)
	}

	if value, ok = os.LookupEnv("ROOT_DIRECTORY"); ok {
		Config.ROOT_DIRECTORY = value
	} else if Config.ENVIRONMENT != development {
		slog.Error("Root directory not set in non-development environment.")
		os.Exit(1)
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			slog.Error("Unable to set root directory.")
			os.Exit(1)
		}
		Config.ROOT_DIRECTORY = cwd
	}

	if value, ok = os.LookupEnv("SECRET_KEY"); ok {
		Config.SECRET_KEY = []byte(value)
	} else if Config.ENVIRONMENT != development {
		slog.Error("Secret key not set in non-development environment.")
		os.Exit(1)
	} else {
		Config.SECRET_KEY = []byte(secret_key)
	}

	if Config.ENVIRONMENT == production {
		if value, ok = os.LookupEnv("LISTEN_PORT"); ok { if intvalue, err := strconv.ParseInt(value, 10, 64); err != nil {
				if !ok || intvalue < 20 || intvalue > 60000 {
					slog.Error("A valid port not set for production environment.")
					os.Exit(1)
				}
			} else {
				Config.LISTEN_PORT = int(intvalue)
			}
		}
	}
}
