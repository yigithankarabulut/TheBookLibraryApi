package bookService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"testing"
)

func TestBookStoreService_FilterByWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := bss.FilterBy(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_FilterByWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		filterErr: errStorageFilter,
	}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	filterRequest := bookService.FilterByBookRequest{Filter: map[string]interface{}{"author": "Yiğit", "language": "Golang"}}

	if _, err := bss.FilterBy(context.Background(), &filterRequest); !errors.Is(
		err,
		errStorageFilter,
	) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_FilterBy(t *testing.T) {
	mockStorage := &mockStorage{bookDb: database.FakeConnectBook()}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	filterRequest := bookService.FilterByBookRequest{Filter: map[string]interface{}{"author": "Yiğit", "language": "Golang"}}

	if _, err := bss.FilterBy(context.Background(), &filterRequest); err != nil {
		t.Error("error occurred")
	}
}
