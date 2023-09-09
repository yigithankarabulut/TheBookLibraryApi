package bookService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"strings"
	"testing"
)

func TestBookStoreService_GetWithCancel(t *testing.T) {
	mockStorage := mockStorage{}
	bss := bookService.New(bookService.WithStorage(&mockStorage))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := bss.Get(ctx, 1); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_GetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		getErr: errStorageGet,
	}
	bss := bookService.New(bookService.WithStorage(mockStorage))
	if _, err := bss.Get(context.Background(), 1); !strings.Contains(
		err.Error(),
		"bookService.Get ",
	) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_Get(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))
	if _, err := bss.Get(context.Background(), 1); err != nil {
		t.Error("error occurred")
	}
}
