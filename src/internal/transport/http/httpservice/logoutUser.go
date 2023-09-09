package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s *httpStorageHandler) LogoutUser(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.Status(200).JSON(fiber.Map{
		"error":   nil,
		"message": "User logged out successfully",
	})
}
