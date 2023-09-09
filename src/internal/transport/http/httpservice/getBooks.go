package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (s *httpStorageHandler) GetBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := s.bookService.List(ctx)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"data":  "empty book list",
		})
	}
	c.Status(200).JSON(fiber.Map{
		"error": nil,
		"data":  result,
	})
	return nil
}
