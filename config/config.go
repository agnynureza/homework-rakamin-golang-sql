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
	JWTSecret   string
	JWTExpired  int
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
		JWTSecret:   GetString("JWT_SECRET_KEY"),
		JWTExpired:  GetInt("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"),
	}

	return appConfig
}

func GetJWTSecret() string {
	return appConfig.JWTSecret
}

func GetJWTExp() int {
	return appConfig.JWTExpired
}
