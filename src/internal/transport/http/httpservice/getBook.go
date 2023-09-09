package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func (s *httpStorageHandler) GetBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, convErr := strconv.Atoi(c.Query("id"))
	if convErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid query parameter",
			"data":  nil,
		})
	}
	result, err := s.bookService.Get(ctx, id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}
	c.Status(200).JSON(fiber.Map{
		"error": nil,
		"data":  result,
	})
	return nil
}
