package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/basehttphandler"
	"log/slog"
	"time"
)

type HTTPService interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	LogoutUser(c *fiber.Ctx) error
	GetBooks(c *fiber.Ctx) error
	GetBook(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
	Router(app fiber.Router)
}

type httpStorageHandler struct {
	basehttphandler.Handler
	userService userService.UserStoreService
	bookService bookService.BookStoreService
}

type StoreHandlerOption func(*httpStorageHandler)

func WithUserService(srvc userService.UserStoreService) StoreHandlerOption {
	return func(s *httpStorageHandler) {
		s.userService = srvc
	}
}

func WithBookService(srvc bookService.BookStoreService) StoreHandlerOption {
	return func(s *httpStorageHandler) {
		s.bookService = srvc
	}
}

func WithContextTimeout(d time.Duration) StoreHandlerOption {
	return func(s *httpStorageHandler) {
		s.Handler.CancelTimeout = d
	}
}

func WithServerEnv(env string) StoreHandlerOption {
	return func(s *httpStorageHandler) {
		s.Handler.ServerEnv = env
	}
}

func WithLogger(l *slog.Logger) StoreHandlerOption {
	return func(s *httpStorageHandler) {
		s.Handler.Logger = l
	}
}

func New(options ...StoreHandlerOption) HTTPService {
	s := &httpStorageHandler{
		Handler: basehttphandler.Handler{},
	}

	for _, option := range options {
		option(s)
	}

	return s
}
