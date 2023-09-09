package httpservice

import (
	"context"
	"github.com/gofiber/fiber/v2"
	bs "github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"net/http"
	"time"
)

func (s *httpStorageHandler) CreateBook(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data": map[string]interface{}{
				"MemberNumber": book.Id,
				"Autor":        book.Author,
				"Title":        book.Title,
				"PublishDate":  book.PublishDate,
			},
		})
	}
	if book.Id == 0 || book.Author == "" || book.Title == "" || book.ISBN == "" || book.PageCount == 0 || book.Category == "" || book.Language == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "bad request",
			"data": map[string]interface{}{
				"Id":        book.Id,
				"Author":    book.Author,
				"Title":     book.Title,
				"Isbn":      book.ISBN,
				"PageCount": book.PageCount,
				"Category":  book.Category,
				"Language":  book.Language,
			},
		})
	}
	result, err := s.bookService.Set(ctx, &bs.SetBookRequest{Book: book})
	if err != nil {
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
