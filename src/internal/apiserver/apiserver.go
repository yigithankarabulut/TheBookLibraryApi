package apiserver

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/httpservice"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ContextCancelTimeout = 5 * time.Second
)

type apiServer struct {
	logLevel  slog.Level
	logger    *slog.Logger
	serverEnv string
}

type Option func(*apiServer)

func WithLogger(l *slog.Logger) Option {
	return func(s *apiServer) {
		s.logger = l
	}
}

func WithServerEnv(env string) Option {
	return func(s *apiServer) {
		s.serverEnv = env
	}
}

func WithLogLevel(level string) Option {
	return func(s *apiServer) {
		var logLevel slog.Level

		switch level {
		case "debug":
			logLevel = slog.LevelDebug
		case "info":
			logLevel = slog.LevelInfo
		case "warn":
			logLevel = slog.LevelWarn
		case "error":
			logLevel = slog.LevelError
		default:
			logLevel = slog.LevelInfo
		}

		s.logLevel = logLevel
	}
}

func New(options ...Option) error {
	apisrv := &apiServer{
		logLevel: slog.LevelInfo,
	}
	usersDb, booksDb := database.Connect()
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	for _, o := range options {
		o(apisrv)
	}

	if apisrv.logger == nil {
		logHandlerOpts := &slog.HandlerOptions{Level: apisrv.logLevel}
		logHandler := slog.NewJSONHandler(os.Stdout, logHandlerOpts)
		apisrv.logger = slog.New(logHandler)
	}
	slog.SetDefault(apisrv.logger)

	if apisrv.serverEnv == "" {
		apisrv.serverEnv = "production"
	}

	logger := apisrv.logger

	userStorage := users.New(users.WithUserDb(usersDb))
	bookStorage := books.New(books.WithBookDb(booksDb))

	usersService := userService.New(userService.WithStorage(userStorage))
	booksService := bookService.New(bookService.WithStorage(bookStorage))

	httpStoreHandler := httpservice.New(
		httpservice.WithUserService(usersService),
		httpservice.WithBookService(booksService),
		httpservice.WithContextTimeout(ContextCancelTimeout),
		httpservice.WithServerEnv(apisrv.serverEnv),
		httpservice.WithLogger(logger),
	)

	shutdown := make(chan os.Signal, 1)
	apiError := make(chan error, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	httpStoreHandler.Router(app)

	go func() {
		logger.Info("starting api server...")
		apiError <- app.Listen(":8080")
	}()

	select {
	case err := <-apiError:
		return fmt.Errorf("listening error: %w", err)
	case <-shutdown:
		logger.Info("starting shutdown", "pid", os.Getpid())
		time.Sleep(1 * time.Second)
		defer logger.Info("shutdown completed", "pid", os.Getpid())
	}
	return nil
}
