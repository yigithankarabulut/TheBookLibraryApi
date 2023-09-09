package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/middlewares"
)

func (r *httpStorageHandler) Router(app fiber.Router) {
	app.Post("/users", r.RegisterUser)
	app.Post("/login", r.LoginUser)

	app.Use(middlewares.JWTMiddleware())

	app.Get("/logout", r.LogoutUser)
	app.Post("/books", r.CreateBook)
	app.Post("/books/filter", r.FilterBooks)
	app.Get("/books", r.GetBooks)
	app.Get("/book/", r.GetBook)       // ?mem_number=123
	app.Put("/book/", r.UpdateBook)    // ?mem_number=123
	app.Delete("/book/", r.DeleteBook) // ?mem_number=123
}