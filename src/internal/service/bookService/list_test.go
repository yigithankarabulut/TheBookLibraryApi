package bookService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"testing"
)

func TestBookStoreService_ListWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := bss.List(ctx); !errors.Is(err, ctx.Err()) {
		t.Error("error occurred")
	}
}

func TestBookStoreService_List(t *testing.T) {
	mockStorage := &mockStorage{bookDb: database.FakeConnectBook()}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	if _, err := bss.List(context.Background()); err != nil {
		t.Error("error occurred")
	}
}
