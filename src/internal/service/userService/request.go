package userService

import "github.com/yigithankarabulut/libraryapi/src/internal/storage/models"

type SetUserRequest struct {
	User models.User
}

type UpdateUserRequest struct {
	MemberNumber int
	UpdateData   map[string]interface{}
}
