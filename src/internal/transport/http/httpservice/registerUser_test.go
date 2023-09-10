package httpservice_test

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHttpStorageHandler_RegisterUser_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
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

func TestHttpStorageHandler_RegisterUser_BodyReadError(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/users", &errorReader{})
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, -1)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestHttpStorageHandler_RegisterUser_BodyUnmarshal(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"username": "test", "password": "test}`))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusBadRequest, res.StatusCode)
	}

}

func TestHttpStorageHandler_RegisterUser_EmptyBody(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	req.Header.Set("Content-Type", "application/json")
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

func TestHttpStorageHandler_RegisterUser_FieldsIsEmpty(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := strings.NewReader(`{"username": "", "password": ""}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
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

func TestHttpStorageHandler_RegisterUser_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1),
		httpservice.WithUserService(&mockUserService{
			userSetErr:   context.DeadlineExceeded,
			userGetErr:   context.DeadlineExceeded,
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 1, FirstName: "test", LastName: "test", Password: "test", Email: "test"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := strings.NewReader(`{
    "member_number":1,
    "firstName":"test",
    "lastName":"test",
    "email":"test",
    "password":"test"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/users", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
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

func TestHttpStorageHandler_RegisterUser_IdExist(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10*time.Second),
		httpservice.WithUserService(&mockUserService{
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 1, FirstName: "test", LastName: "test", Password: "test", Email: "test"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := strings.NewReader(`{
    "member_number":1,
    "firstName":"test",
    "lastName":"test",
    "email":"test",
    "password":"test"
	}`)
	req := httptest.NewRequest(http.MethodPost, "/users", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusBadGateway {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusBadGateway, res.StatusCode)
	}

	shouldContain := "user already exist"
	resBody, _ := io.ReadAll(res.Body)

	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}

}

func TestHttpStorageHandler_RegisterUser_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10*time.Second),
		httpservice.WithUserService(&mockUserService{
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 1, FirstName: "test", LastName: "test", Password: "test", Email: "test"}},
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := strings.NewReader(`{
	"member_number":2,
	"firstName":"test",
	"lastName":"test",
	"email":"test",
	"password":"test"
	}`)
	req := httptest.NewRequest(http.MethodPost, "/users", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusOK, res.StatusCode)
	}

	shouldContain := "user created successfully"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}

}
