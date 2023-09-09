package userService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"testing"
)

func TestUserStoreService_ListWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := uss.List(ctx); !errors.Is(err, ctx.Err()) {
		t.Error("error occurred")
	}
}

func TestUserStoreService_List(t *testing.T) {
	mockStorage := &mockStorage{userDb: database.FakeConnectUser()}
	uss := userService.New(userService.WithStorage(mockStorage))

	if _, err := uss.List(context.Background()); err != nil {
		t.Error("error occurred")
	}
}
