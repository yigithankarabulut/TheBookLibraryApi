package bookService

import (
	"context"
	"fmt"
)

func (s *bookStoreService) Delete(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.storage.Delete(id); err != nil {
			return fmt.Errorf("bookService.Delete err: %w", err)
		}
		return nil
	}
}
