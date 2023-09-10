package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s *httpStorageHandler) LogoutUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return c.Status(504).SendString("context deadline exceeded")
	default:
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})
		return c.Status(200).JSON(fiber.Map{
			"message": "User logged out successfully",
		})
	}
}
