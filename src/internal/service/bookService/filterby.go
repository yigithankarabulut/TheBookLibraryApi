package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) FilterBy(ctx context.Context, req *FilterByBookRequest) (*ListBooksResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		items, err := s.storage.FilterBy(req.Filter)
		if err != nil {
			return nil, fmt.Errorf("bookService.FilterBy err: %w", err)
		}

		response := make(ListBooksResponse, len(items))
		var i int
		for _, v := range items {
			response[i].Book = v
			i++
		}
		return &response, nil
	}
}
