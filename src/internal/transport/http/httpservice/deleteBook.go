package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func (s *httpStorageHandler) DeleteBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data":  "invalid query parameter",
		})
	}
	resErr := s.bookService.Delete(ctx, id)
	if resErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": resErr.Error(),
			"data": map[string]interface{}{
				"id": id,
			},
		})
	}
	c.Status(200).JSON(fiber.Map{
		"error": nil,
		"data":  "deleted successfully",
	})
	return nil
}
