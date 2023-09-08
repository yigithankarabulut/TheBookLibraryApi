package userService

import (
	"context"
	"fmt"
)

func (s *userStoreService) Get(ctx context.Context, memberNumber int) (*UserResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		user, err := s.storage.Get(memberNumber)
		if err != nil {
			return nil, fmt.Errorf("userservice.Get err: %w", err)
		}
		return &UserResponse{
			User: user,
		}, nil
	}
}
