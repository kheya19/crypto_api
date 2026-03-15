package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/crypto_api/crypto"
	"github.com/kheya19/crypto_api/model"
)

func Encrypt(c *fiber.Ctx) error {
	var req model.EncryptRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if req.PlainText == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "PlainText is required",
		})
	}
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	cipherText, err := crypto.Encrypt(req.PlainText, key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Encryption failed",
		})
	}
	return c.JSON(model.EncryptResponse{CipherText: cipherText})
}
