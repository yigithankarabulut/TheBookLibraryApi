package userService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"testing"
)

func TestUserStoreService_UpdateWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := uss.Update(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestUserStoreService_UpdateWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		updateErr: errStorageUpdate,
	}
	uss := userService.New(userService.WithStorage(mockStorage))

	updateRequest := userService.UpdateUserRequest{MemberNumber: 10, UpdateData: map[string]interface{}{
		"firstName": "firstName",
		"lastName":  "lastName",
		"email":     "email",
		"password":  "password",
	}}

	if _, err := uss.Update(context.Background(), &updateRequest); !errors.Is(
		err,
		errStorageUpdate,
	) {
		t.Error("error not occurred")
	}
}

func TestUserStoreService_Update(t *testing.T) {
	mockStorage := &mockStorage{
		userDb: database.FakeConnectUser(),
	}
	uss := userService.New(userService.WithStorage(mockStorage))
	updateRequest := userService.UpdateUserRequest{MemberNumber: 10, UpdateData: map[string]interface{}{
		"firstName": "firstName",
		"lastName":  "lastName",
		"email":     "email",
		"password":  "password",
	}}
	if _, err := uss.Update(context.Background(), &updateRequest); err != nil {
		t.Error("error occurred")
	}
}
