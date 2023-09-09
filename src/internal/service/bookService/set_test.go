package bookService_test

import (
	"context"
	"errors"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"testing"
)

func TestBookStoreService_SetWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := bss.Set(ctx, nil); !errors.Is(err, ctx.Err()) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_SetWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{
		setErr: errStorageSet,
	}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	setRequest := bookService.SetBookRequest{Book: models.Book{Id: 3, Title: "title", Author: "author", Language: "Golang", Category: "Programming"}}
	if _, err := bss.Set(context.Background(), &setRequest); !errors.Is(
		err,
		errStorageSet,
	) {
		t.Error("error not occurred")
	}
}

func TestBookStoreService_Set(t *testing.T) {
	mockStorage := &mockStorage{bookDb: database.FakeConnectBook()}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	setRequest := bookService.SetBookRequest{Book: models.Book{Id: 3, Title: "title", Author: "author", Language: "Golang", Category: "Programming"}}

	if _, err := bss.Set(context.Background(), &setRequest); err != nil {
		t.Error("error occurred")
	}
}
