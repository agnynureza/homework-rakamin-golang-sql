package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount := GetInt("SERVER_READ_TIMEOUT")

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
