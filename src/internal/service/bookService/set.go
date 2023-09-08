package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) Set(ctx context.Context, sr *SetBookRequest) (*BookResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		book, err := s.storage.Set(sr.Book)
		if err != nil {
			return nil, fmt.Errorf("bookservice.Set err %w", err)
		}
		return &BookResponse{
			Book: book,
		}, nil
	}
}
