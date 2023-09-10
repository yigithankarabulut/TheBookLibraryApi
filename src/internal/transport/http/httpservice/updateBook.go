package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"strconv"
)

func (s *httpStorageHandler) UpdateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	var bookUpdateResult bookService.UpdateBookRequest
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("invalid query params")
	}
	bookUpdateResult.Id = id
	bookUpdateResult.UpdateData = make(map[string]interface{})
	if err = c.BodyParser(&bookUpdateResult.UpdateData); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusInternalServerError).SendString("empty body/payload")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("invalid body/payload")
	}
	if getRes, err := s.bookService.Get(ctx, id); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(fiber.StatusGatewayTimeout).SendString("context deadline exceeded")
		}
	} else if getRes.Book.Id != bookUpdateResult.Id {
		return c.Status(fiber.StatusNotFound).SendString("book not found")
	}
	_, err2 := s.bookService.Update(ctx, &bookUpdateResult)
	if err2 != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(504).SendString("context deadline exceeded")
		}
		return c.Status(400).JSON(fiber.Map{
			"error": err2.Error(),
			"data": map[string]interface{}{
				"id": bookUpdateResult.Id,
			},
		})
	}
	c.Status(200).SendString("book updated")
	return nil
}
