package handlers

import (
	"github.com/agnynureza/homework-rakamin-golang-sql/common/utils"
	"github.com/gofiber/fiber/v2"
)

type TokenHandler struct{}

type TokenHandlerInterface interface {
	GetNewAccessToken(c *fiber.Ctx) error
}

func (t *TokenHandler) GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error":        false,
		"msg":          "success create token",
		"access_token": token,
	})
}
