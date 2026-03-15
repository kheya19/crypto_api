package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/crypto_api/crypto"
	"github.com/kheya19/crypto_api/model"
)

func Decrypt(c *fiber.Ctx) error {
	var req model.DecryptRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if req.CipherText == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "CipherText is required",
		})
	}
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	plainText, err := crypto.Decrypt(req.CipherText, key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Decryption failed",
		})
	}
	return c.JSON(model.DecryptResponse{PlainText: plainText})
}
