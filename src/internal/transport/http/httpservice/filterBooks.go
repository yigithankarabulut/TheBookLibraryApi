package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"time"
)

func (s *httpStorageHandler) FilterBooks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var filter bookService.FilterByBookRequest
	if err := c.BodyParser(&filter.Filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}
	result, err := s.bookService.FilterBy(ctx, &filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot find books with given filter",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   "null",
		"message": "books filtered successfully",
		"data":    result,
	})
}
