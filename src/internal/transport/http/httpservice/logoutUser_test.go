package httpservice_test

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestHttpStorageHandler_LogoutUser_InvalidMethod(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New()
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
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

func TestHttpStorageHandler_LogoutUser_Timeout(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(-1),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
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

func TestHttpStorageHandler_LogoutUser_Success(t *testing.T) {
	app := fiber.New()
	handler := httpservice.New(
		httpservice.WithContextTimeout(10 * time.Second),
	)
	r := fiber.Router(app)
	handler.Router(r)

	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
	cookie := http.Cookie{
		Name:  "jwt",
		Value: "test",
	}
	req.AddCookie(&cookie)
	res, err := app.Test(req)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if res.StatusCode != fiber.StatusOK {
		t.Errorf("wrong status code, want: %d, got: %d", fiber.StatusOK, res.StatusCode)
	}

	shouldContain := fmt.Sprintf("{\"message\":\"User logged out successfully\"}")
	resBody, _ := io.ReadAll(res.Body)

	if !strings.Contains(string(resBody), shouldContain) {
		t.Errorf("wrong body message, want: %s, got: %s", shouldContain, string(resBody))
	}
}
