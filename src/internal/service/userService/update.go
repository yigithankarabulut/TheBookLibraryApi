package userService

import (
	"context"
	"fmt"
)

func (s *userStoreService) Update(ctx context.Context, ur *UpdateUserRequest) (*UserResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		user, err := s.storage.Update(ur.MemberNumber, ur.UpdateData)
		if err != nil {
			return nil, fmt.Errorf("userservice.Update err: %w", err)
		}
		return &UserResponse{
			User: user,
		}, nil
	}
}
