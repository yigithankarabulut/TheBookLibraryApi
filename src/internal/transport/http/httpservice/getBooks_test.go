package httpservice_test

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestHttpStorageHandler_GetBooks_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodDelete, "/books", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusMethodNotAllowed {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusMethodNotAllowed, res.StatusCode)
	}

	shouldContain := "Method Not Allowed"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_GetBooks_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(time.Second*-1),
		httpservice.WithBookService(&mockBookService{
			bookListErr: context.DeadlineExceeded,
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != http.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusGatewayTimeout, res.StatusCode)
	}

	shouldContain := "context deadline exceeded"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_GetBooks_EmptyList(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithBookService(&mockBookService{}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestHttpStorageHandler_GetBooks_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithBookService(&mockBookService{bookListResponse: &bookService.ListBooksResponse{bookService.BookResponse{Book: models.Book{Id: 99, Category: "test", Language: "test", PageCount: 999, Author: "test", Title: "test", ISBN: "9999", Description: "test"}}}}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/books", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", http.StatusOK, res.StatusCode)
	}

	shouldEqual := `{"data":[{"Book":{"id":99,"title":"test","author":"test","publishDate":"0001-01-01T00:00:00Z","isbn":"9999","pageCount":999,"category":"test","language":"test","description":"test"}}]}`
	resBody, _ := io.ReadAll(res.Body)
	if string(resBody) != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, string(resBody))
	}
}
