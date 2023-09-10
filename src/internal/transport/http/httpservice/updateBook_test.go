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

func TestHttpStorageHandler_UpdateBook_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/updbook", nil)
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

func TestHttpStorageHandler_UpdateBook_BodyReadError(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPut, "/updbook", &errorReader{})
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, -1)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestHttpStorageHandler_UpdateBook_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1),
		httpservice.WithBookService(&mockBookService{
			bookResponse:  &bookService.BookResponse{Book: models.Book{Id: 1, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}},
			bookUpdateErr: context.DeadlineExceeded,
			bookGetErr:    context.DeadlineExceeded,
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := `{
    "title":"Postmans",
    "author":"YK",
    "category":"rest",
    "pageCount":750
	}`
	req, _ := http.NewRequest(http.MethodPut, "/updbook/?id=1", strings.NewReader(handlerRequest))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusGatewayTimeout, res.StatusCode)
	}

	shouldContain := "context deadline exceeded"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}

}

func TestHttpStorageHandler_UpdateBook_InvalidQueryParams(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPut, "/updbook", nil)
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

func TestHttpStorageHandler_UpdateBook_BodyUnmarshal(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPut, "/updbook/?id=1", strings.NewReader(`{
    "title":"Postmans",
    "author":"YK",
    "category":"rest}`))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusInternalServerError, res.StatusCode)
	}
}

func TestHttpStorageHandler_UpdateBook_IdNotFound(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10*time.Second),
		httpservice.WithBookService(&mockBookService{
			bookResponse: &bookService.BookResponse{Book: models.Book{Id: 1, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := `{
    "title":"test",
    "author":"test",
    "category":"test",
    "pageCount":750
	}`

	req, _ := http.NewRequest(http.MethodPut, "/updbook/?id=3", strings.NewReader(handlerRequest))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusNotFound, res.StatusCode)
	}

	shouldContain := "book not found"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_UpdateBook_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10*time.Second),
		httpservice.WithBookService(&mockBookService{
			bookResponse: &bookService.BookResponse{Book: models.Book{Id: 1, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := `{
    "title":"Postmans",
    "author":"YK",
    "category":"rest",
    "pageCount":750
	}`

	req, _ := http.NewRequest(http.MethodPut, "/updbook/?id=1", strings.NewReader(handlerRequest))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusOK, res.StatusCode)
	}

	shouldContain := "book updated"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}
