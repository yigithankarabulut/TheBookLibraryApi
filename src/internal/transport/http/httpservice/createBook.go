package httpservice

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	bs "github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"net/http"
)

func (s *httpStorageHandler) CreateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.Handler.CancelTimeout)
	defer cancel()
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "empty body/payload",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if book.Id == 0 || book.Author == "" || book.Title == "" || book.ISBN == "" || book.PageCount == 0 || book.Category == "" || book.Language == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty field/fields",
		})
	}

	result, err := s.bookService.Set(ctx, &bs.SetBookRequest{Book: book})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return c.Status(504).SendString("context deadline exceeded")
		}
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
			"data": map[string]interface{}{
				"id":   book.Id,
				"book": "not exist",
			},
		})
	}
	c.Status(200).JSON(fiber.Map{
		"error": nil,
		"data":  result.Book,
	})
	return nil
}
