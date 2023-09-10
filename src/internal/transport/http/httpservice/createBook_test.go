package httpservice_test

import (
	"bytes"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type errorReader struct{}

func (e *errorReader) Read(_ []byte) (n int, err error) {
	return 0, errors.New("forced error")
}

func TestHttpStorageHandler_CreateBook_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/crtbooks", nil)
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

func TestHttpStorageHandler_CreateBook_BodyReadError(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/crtbooks", &errorReader{})
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	_, err := app.Test(req, -1)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestHttpStorageHandler_CreateBook_BodyUnmarshal(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{"key": "key", "value": "123}`)
	req, _ := http.NewRequest(http.MethodPost, "/crtbooks", handlerRequest)
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

func TestHttpStorageHandler_CreateBook_EmptyBody(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/crtbooks", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusBadRequest, res.StatusCode)
	}

	shouldContain := "empty body/payload"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_CreateBook_FieldsIsEmpty(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{"id":0,"title":""}`)

	req, _ := http.NewRequest(http.MethodPost, "/crtbooks", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusBadRequest {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusBadRequest, res.StatusCode)
	}

	shouldContain := "empty field/fields"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_CreateBook_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1),
		httpservice.WithBookService(&mockBookService{
			bookSetErr: context.DeadlineExceeded,
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	jsonInput := `{
        "id": 1112,
        "title": "Sample Book",
        "author": "John Doe",
        "isbn": "1234567890",
        "pageCount": 200,
        "category": "Fiction",
        "language": "English"
    }`

	req := httptest.NewRequest(http.MethodPost, "/crtbooks", strings.NewReader(jsonInput))
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

func TestHttpStorageHandler_CreateBook_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10*time.Second),
		httpservice.WithBookService(&mockBookService{
			bookResponse: &bookService.BookResponse{Book: models.Book{Id: 171, Title: "PostmanPostBookRequest", Author: "Yigithan", ISBN: "112233445566", PageCount: 250, Category: "rest", Language: "kotlin"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	jsonInput := `{
    "id":17,
    "title":"PostmanPostBookRequest",
    "author":"Yigithan",
    "isbn":"112233445566",
    "pageCount":250,
    "category":"rest",
    "language":"kotlin"
	}`
	req := httptest.NewRequest(http.MethodPost, "/crtbooks", strings.NewReader(jsonInput))
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

	shouldEqual := `{"data":{"id":171,"title":"PostmanPostBookRequest","author":"Yigithan","publishDate":"0001-01-01T00:00:00Z","isbn":"112233445566","pageCount":250,"category":"rest","language":"kotlin"},"error":null}`
	resBody, _ := io.ReadAll(res.Body)
	if string(resBody) != shouldEqual {
		t.Errorf("wrong body message, want: %s, got: %s", shouldEqual, string(resBody))
	}
}
