package bookService

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
)

type BookStoreService interface {
	Set(context.Context, *SetBookRequest) (*BookResponse, error)
	Get(context.Context, int) (*BookResponse, error)
	Update(context.Context, *UpdateBookRequest) (*BookResponse, error)
	Delete(context.Context, int) error
	List(context.Context) (*ListBooksResponse, error)
	FilterBy(context.Context, *FilterByBookRequest) (*ListBooksResponse, error)
}

type bookStoreService struct {
	storage books.BookStorer
}

type BookStoreServiceOption func(*bookStoreService)

func WithStorage(strg books.BookStorer) BookStoreServiceOption {
	return func(s *bookStoreService) {
		s.storage = strg
	}
}

func New(options ...BookStoreServiceOption) BookStoreService {
	bss := &bookStoreService{}

	for _, o := range options {
		o(bss)
	}

	return bss
}
