package userService

import (
	"context"
	"fmt"
)

func (s *userStoreService) Set(ctx context.Context, sr *SetUserRequest) (*UserResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		user, err := s.storage.Set(sr.User)
		if err != nil {
			return nil, fmt.Errorf("userservice.Set err: %w", err)
		}
		return &UserResponse{
			User: user,
		}, nil
	}
}
