package bookService

import "github.com/yigithankarabulut/libraryapi/src/internal/storage/models"

type BookResponse struct {
	Book models.Book
}

type ListBooksResponse []BookResponse
