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

func TestHttpStorageHandler_GetBook_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)
	req, _ := http.NewRequest(http.MethodPatch, "/book", nil)
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

func TestHttpStorageHandler_GetBook_InvalidQueryParams(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/book", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusNotFound, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	shouldContain := "invalid query params"
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_GetBook_ParamKeyNotFound(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/book/?foo=test", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusNotFound, res.StatusCode)
	}
	resBody, _ := io.ReadAll(res.Body)
	shouldContain := "invalid query params"
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_GetBook_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1*time.Second),
		httpservice.WithBookService(&mockBookService{
			bookGetErr: context.DeadlineExceeded,
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/book/?id=1", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusGatewayTimeout, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	shouldContain := "context deadline exceeded"
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_GetBook_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithBookService(&mockBookService{
			bookResponse: &bookService.BookResponse{Book: models.Book{Id: 99, Category: "test", Language: "test", PageCount: 999, Author: "test", Title: "test", ISBN: "9999", Description: "test"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/book/?id=99", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusOK, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	shouldEqual := `{"data":{"Book":{"id":99,"title":"test","author":"test","publishDate":"0001-01-01T00:00:00Z","isbn":"9999","pageCount":999,"category":"test","language":"test","description":"test"}}}`
	if string(resBody) != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, string(resBody))
	}
}
