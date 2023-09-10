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

func TestHttpStorageHandler_FilterBooks_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/books/filter", nil)
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

func TestHttpStorageHandler_FilterBooks_BodyReadError(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/books/filter", &errorReader{})
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)

	_, err := app.Test(req)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestHttpStorageHandler_FilterBooks_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithBookService(&mockBookService{bookFilterErr: context.DeadlineExceeded}),
		httpservice.WithContextTimeout(-1),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/books/filter/?id=1", strings.NewReader(`{"key": "key", "value": "value"}`))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)

	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusGatewayTimeout {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusInternalServerError, res.StatusCode)
	}

	shouldContain := "context deadline exceeded"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_FilterBooks_EmptyBody(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/books/filter", nil)
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusNotFound, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	shouldContain := "empty body/payload"
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_FilterBooks_BodyUnmarshall(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerReq := `{"key": "key", "value": "value}`
	req, _ := http.NewRequest(http.MethodPost, "/books/filter", strings.NewReader(handlerReq))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusBadRequest, res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	shouldContain := "invalid body/payload"
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_FilterBooks_Success(t *testing.T) {
	app := fiber.New()
	book1 := bookService.BookResponse{Book: models.Book{Id: 171, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}}
	book2 := bookService.BookResponse{Book: models.Book{Id: 172, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}}
	bookResponse := new(bookService.ListBooksResponse)
	for i := 0; i < 2; i++ {
		*bookResponse = append(*bookResponse, book1, book2)
	}
	handler := httpservice.New(
		httpservice.WithBookService(&mockBookService{bookListResponse: bookResponse}),
		httpservice.WithContextTimeout(10*time.Second),
	)

	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := `{"title":"PostmanPostBookRequest","author":"Yigithan"}`

	req, _ := http.NewRequest(http.MethodPost, "/books/filter", strings.NewReader(handlerRequest))
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

}
