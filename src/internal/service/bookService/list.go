package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) List(ctx context.Context) (*ListBooksResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		books, err := s.storage.List()
		if err != nil {
			return nil, fmt.Errorf("bookService.List err %w", err)
		}
		BooksListResult := make(ListBooksResponse, len(books))
		var i int
		for _, book := range books {
			BooksListResult[i].Book = book
			i++
		}
		return &BooksListResult, nil
	}
}
