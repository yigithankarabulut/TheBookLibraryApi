package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yigithankarabulut/libraryapi/src/internal/transport/http/middlewares"
)

func (r *httpStorageHandler) Router(app fiber.Router) {
	app.Post("/users", r.RegisterUser)
	app.Post("/login", r.LoginUser)

	app.Use(middlewares.JWTMiddleware())

	app.Get("/logout", r.LogoutUser)
	app.Post("/crtbooks", r.CreateBook)
	app.Post("/books/filter", r.FilterBooks)
	app.Get("/books", r.GetBooks)
	app.Get("/book/", r.GetBook)          // ?id=123
	app.Put("/updbook/", r.UpdateBook)    // ?id=123
	app.Delete("/delbook/", r.DeleteBook) // ?id=123
}
