package bookService

import "github.com/yigithankarabulut/libraryapi/src/internal/storage/models"

type SetBookRequest struct {
	Book models.Book
}

type UpdateBookRequest struct {
	Id         int
	UpdateData map[string]interface{}
}

type FilterByBookRequest struct {
	Filter map[string]interface{}
}
