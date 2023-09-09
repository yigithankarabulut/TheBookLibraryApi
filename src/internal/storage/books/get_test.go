package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"testing"
)

func TestBookStorage_Get(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "TestBookStorage_Get", args: args{id: 0}, wantErr: true},
		{name: "TestBookStorage_Get", args: args{id: 1}, wantErr: false},
		{name: "TestBookStorage_Get", args: args{id: 3}, wantErr: false},
		{name: "TestBookStorage_Get", args: args{id: 15}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := books.New(books.WithBookDb(fakeDb))
			if _, err := usr.Get(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
