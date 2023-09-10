package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (s *httpStorageHandler) GetBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	id, convErr := strconv.Atoi(c.Query("id"))
	if convErr != nil {
		return c.Status(fiber.StatusNotFound).SendString("invalid query params")
	}
	result, err := s.bookService.Get(ctx, id)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(504).SendString("context deadline exceeded")
		}
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}
	c.Status(200).JSON(fiber.Map{
		"data": result,
	})
	return nil
}
