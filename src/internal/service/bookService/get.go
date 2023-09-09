package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) Get(ctx context.Context, id int) (*BookResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		result, err := s.storage.Get(id)
		if err != nil {
			return nil, fmt.Errorf("bookService.Get err %w", err)
		}
		return &BookResponse{
			Book: result,
		}, nil
	}
}
