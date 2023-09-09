package bookService_test

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"strings"
	"testing"
	"time"
)

func TestBookStoreService_DeleteWithCancel(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	id := 1
	if err := bss.Delete(ctx, id); err == nil {
		t.Errorf("Delete() error = %v, wantErr %v", err, true)
	}
}

func TestBookStoreService_DeleteWithStorageError(t *testing.T) {
	mockStorage := &mockStorage{deleteErr: errStorageDelete, bookDb: database.FakeConnectBook()}
	bss := bookService.New(bookService.WithStorage(mockStorage))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := bss.Delete(ctx, 0); !strings.Contains(
		err.Error(),
		"bookService.Delete ",
	) {
		t.Errorf("Delete() error = %v, wantErr %v", err, true)
	}
}

func TestBookStoreService_Delete(t *testing.T) {
	mockStorage := &mockStorage{}
	bss := bookService.New(bookService.WithStorage(mockStorage))

	if err := bss.Delete(context.Background(), 1); err != nil {
		t.Errorf("Delete() error = %v, wantErr %v", err, false)
	}
}
