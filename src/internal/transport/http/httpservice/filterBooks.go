package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
)

func (s *httpStorageHandler) FilterBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	var filter bookService.FilterByBookRequest
	if err := c.BodyParser(&filter.Filter); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("empty body/payload")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("invalid body/payload")
	}
	result, err := s.bookService.FilterBy(ctx, &filter)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(504).SendString("context deadline exceeded")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot find books with given filter",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "books filtered successfully",
		"data":    result,
	})
}
