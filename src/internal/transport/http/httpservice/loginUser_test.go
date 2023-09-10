package httpservice_test

import (
	"bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestHttpStorageHandler_LoginUser_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/login", nil)
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

func TestHttpStorageHandler_LoginUser_BodyReadError(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/login", &errorReader{})
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, -1)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

}

func TestHttpStorageHandler_LoginUser_BodyUnmarshal(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":12,
    "password":"ykarabul
	}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusInternalServerError, res.StatusCode)
	}
}

func TestHttpStorageHandler_LoginUser_EmptyBody(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/login", nil)
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

func TestHttpStorageHandler_LoginUser_FieldsIsEmpty(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":12,
    "password":""
	}`)
	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
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

func TestHttpStorageHandler_LoginUser_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1),
		httpservice.WithUserService(&mockUserService{
			userGetErr: context.DeadlineExceeded,
		}),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":12,
    "password":"ykarabul"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
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

func TestHttpStorageHandler_LoginUser_PwdHashFail(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithUserService(&mockUserService{
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 12, Password: "$2a$10$i3rfOE"}}},
		),
		httpservice.WithContextTimeout(10*time.Second),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":12,
    "password":"ykarabul"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusUnauthorized, res.StatusCode)
	}

	shouldContain := "password is wrong"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_LoginUser_IdNotFound(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithUserService(&mockUserService{
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 12, Password: "123456"}},
		}),
		httpservice.WithContextTimeout(10*time.Second),
	)
	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":99999,
    "password":"ykarabul"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusNotFound {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusNotFound, res.StatusCode)
	}

	shouldContain := "user not found"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}

func TestHttpStorageHandler_LoginUser_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithUserService(&mockUserService{
			userResponse: &userService.UserResponse{User: models.User{MemberNumber: 12, Password: "$2a$10$i3rfOE2pLcdZKbY2pSdTpuohvqdJ0IlmKQ39lx/MUMkUBneNtgs2u"}}}),
		httpservice.WithContextTimeout(10*time.Second),
	)

	r := fiber.Router(app)
	handler.Router(r)

	handlerRequest := bytes.NewBufferString(`{
    "member_number":12,
    "password":"ykarabul"
	}`)

	req, _ := http.NewRequest(http.MethodPost, "/login", handlerRequest)
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	if res.StatusCode != fiber.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusOK, res.StatusCode)
	}

	shouldContain := "{\"id\":12,\"message\":\"User logged in successfully\",\"password\":\"ykarabul\"}"
	resBody, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}

}
