package userService

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
)

type UserStoreService interface {
	Set(context.Context, *SetUserRequest) (*UserResponse, error)
	Get(context.Context, int) (*UserResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UserResponse, error)
	Delete(context.Context, int) error
	List(context.Context) (*ListUserResponse, error)
}

type userStoreService struct {
	storage users.UserStorer
}

type UserStoreServiceOption func(*userStoreService)

func WithStorage(strg users.UserStorer) UserStoreServiceOption {
	return func(s *userStoreService) {
		s.storage = strg
	}
}

func New(options ...UserStoreServiceOption) UserStoreService {
	uss := &userStoreService{}

	for _, o := range options {
		o(uss)
	}

	return uss
}
