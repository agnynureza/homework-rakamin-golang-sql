package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

var appConfig *Config

type Config struct {
	AppName     string
	AppPort     string
	LogLevel    string
	Environment string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig = &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetString("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
	}

	return appConfig
}
