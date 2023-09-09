package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestBookStorage_FilterBy(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		filter map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestBookStorage_FilterBy", fields: fields{db: fakeDb}, args: args{filter: map[string]interface{}{"title": "", "author": "", "isbn": ""}}, wantErr: true},
		{name: "TestBookStorage_FilterBy", fields: fields{db: fakeDb}, args: args{filter: map[string]interface{}{"title": "Test", "author": "Test", "isbn": ""}}, wantErr: false},
		{name: "TestBookStorage_FilterBy", fields: fields{db: fakeDb}, args: args{filter: map[string]interface{}{"title": "Test", "author": "Test", "isbn": "Test"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := books.New(books.WithBookDb(tt.fields.db))
			if _, err := usr.FilterBy(tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("FilterBy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
