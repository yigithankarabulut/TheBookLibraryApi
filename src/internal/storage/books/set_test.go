package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestBookStorage_Set(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		book models.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Book
		wantErr bool
	}{
		{name: "TestBookStorage_Set", fields: fields{db: fakeDb}, args: args{book: models.Book{}}, want: models.Book{}, wantErr: true},
		{name: "TestBookStorage_Set", fields: fields{db: fakeDb}, args: args{book: models.Book{Id: 2, Title: "Test", Author: "Test", ISBN: ""}}, want: models.Book{Id: 2, Title: "Test", Author: "Test", ISBN: ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := books.New(books.WithBookDb(tt.fields.db))
			got, err := usr.Set(tt.args.book)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Id != tt.want.Id {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}
