package userService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"testing"
)

func TestUserStoreService_SetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := uss.Set(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestUserStoreService_SetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		setErr: errStorageSet,
	}
	uss := userService.New(userService.WithStorage(mockStorage))

	setRequest := userService.SetUserRequest{User: models.User{MemberNumber: 10, Email: "yk@gmail.com", FirstName: "Yiğit", LastName: "Karabulut", Password: "11223344"}}
	if _, err := uss.Set(context.Background(), &setRequest); !errors.Is(
		err,
		errStorageSet,
	) {
		t.Error("error not occurred")
	}
}

func TestUserStoreService_Set(t *testing.T) {
	mockStorage := &mockStorage{userDb: database.FakeConnectUser()}
	uss := userService.New(userService.WithStorage(mockStorage))

	setRequest := userService.SetUserRequest{User: models.User{MemberNumber: 10, Email: "yk@gmail.com", FirstName: "Yiğit", LastName: "Karabulut", Password: "11223344"}}

	if _, err := uss.Set(context.Background(), &setRequest); err != nil {
		t.Error("error occurred")
	}
}
