package userService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"strings"
	"testing"
)

func TestUserStoreService_GetWithCancel(t *testing.T) {
	mockStorage := mockStorage{}
	uss := userService.New(userService.WithStorage(&mockStorage))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := uss.Get(ctx, 1); !errors.Is(err, ctx.Err()) {
		t.Error("error occurred")
	}
}

func TestUserStoreService_GetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		getErr: errStorageGet,
	}
	uss := userService.New(userService.WithStorage(mockStorage))
	if _, err := uss.Get(context.Background(), 1); !strings.Contains(
		err.Error(),
		"userService.Get ",
	) {
		t.Error("error not occurred")
	}
}

func TestUserStoreService_Get(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))
	if _, err := uss.Get(context.Background(), 1); err != nil {
		t.Error("error occurred")
	}
}
