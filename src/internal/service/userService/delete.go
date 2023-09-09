package userService

import (
	"context"
	"fmt"
)

func (s *userStoreService) Delete(ctx context.Context, memberNumber int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := s.storage.Delete(memberNumber); err != nil {
			return fmt.Errorf("userService.Delete err: %w", err)
		}
		return nil
	}
}
