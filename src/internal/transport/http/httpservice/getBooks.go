package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (s *httpStorageHandler) GetBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	result, err := s.bookService.List(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(504).SendString("context deadline exceeded")
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	} else if result == nil {
		return c.Status(fiber.StatusNotFound).SendString("no data found")
	}
	c.Status(200).JSON(fiber.Map{
		"data": result,
	})
	return nil
}
