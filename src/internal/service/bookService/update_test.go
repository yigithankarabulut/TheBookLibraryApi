package bookService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"strings"
	"testing"
)

func TestBookStoreService_UpdateWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := bss.Update(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_UpdateWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{updateErr: errStorageUpdate}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	updateRequest := bookService.UpdateBookRequest{Id: 1, UpdateData: map[string]interface{}{"title": "title", "author": "author", "language": "Golang", "category": "Programming"}}

	if _, err := bss.Update(context.Background(), &updateRequest); !strings.Contains(
		err.Error(),
		"bookService.Update ",
	) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_Update(t *testing.T) {
	mockStorage := &mockStorage{
		bookDb: database.FakeConnectBook(),
	}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	updateRequest := bookService.UpdateBookRequest{Id: 1, UpdateData: map[string]interface{}{"title": "title", "author": "author", "language": "Golang", "category": "Programming"}}

	if _, err := bss.Update(context.Background(), &updateRequest); err != nil {
		t.Error("error occurred")
	}
}
