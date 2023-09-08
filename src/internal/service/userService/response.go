package userService

import "github.com/yigithankarabulut/libraryapi/src/internal/storage/models"

type UserResponse struct {
	User models.User
}

type ListUserResponse []UserResponse
