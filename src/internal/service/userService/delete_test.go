package userService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"strings"
	"testing"
	"time"
)

func TestUserStoreService_DeleteWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if err := uss.Delete(ctx, 0); !errors.Is(err, ctx.Err()) {
		t.Errorf("Delete() error = %v, wantErr %v", err, true)
	}
}

func TestUserStoreService_DeleteWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{deleteErr: errStorageDelete, userDb: database.FakeConnectUser()}
	uss := userService.New(userService.WithStorage(mockStorage))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := uss.Delete(ctx, 1); !strings.Contains(
		err.Error(),
		"userService.Delete ",
	) {
		t.Errorf("Delete() error = %v, wantErr %v", err, true)
	}
}

func TestUserStoreService_Delete(t *testing.T) {
	mockStorage := &mockStorage{}
	uss := userService.New(userService.WithStorage(mockStorage))

	if err := uss.Delete(context.Background(), 1); err != nil {
		t.Errorf("Delete() error = %v, wantErr %v", err, false)
	}
}
