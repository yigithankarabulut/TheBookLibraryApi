package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"testing"
)

func TestBookStorage_List(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	us := books.New(books.WithBookDb(fakeDb))
	if _, err := us.List(); err != nil {
		t.Errorf("List() error = %v", err)
	}
}
