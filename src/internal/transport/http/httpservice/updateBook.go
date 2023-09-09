package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"strconv"
	"time"
)

func (s *httpStorageHandler) UpdateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bookUpdateResult bookService.UpdateBookRequest
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data":  "invalid query parameter",
		})
	}
	bookUpdateResult.Id = id
	bookUpdateResult.UpdateData = make(map[string]interface{})
	if err = c.BodyParser(&bookUpdateResult.UpdateData); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data": map[string]interface{}{
				"id": bookUpdateResult.Id,
			},
		})
	}
	result, err := s.bookService.Update(ctx, &bookUpdateResult)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
			"data": map[string]interface{}{
				"id": bookUpdateResult.Id,
			},
		})
	}
	c.Status(200).JSON(fiber.Map{
		"error": nil,
		"data":  result.Book,
	})
	return nil
}
