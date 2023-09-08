package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) Update(ctx context.Context, sr *UpdateBookRequest) (*BookResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		book, err := s.storage.Update(sr.Id, sr.UpdateData)
		if err != nil {
			return nil, fmt.Errorf("bookservice.Update err %w", err)
		}
		return &BookResponse{
			Book: book,
		}, nil
	}
}
