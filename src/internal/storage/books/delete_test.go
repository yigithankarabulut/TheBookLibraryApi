package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestBookStorage_Delete(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestBookStorage_Delete", fields: fields{db: fakeDb}, args: args{id: 0}, wantErr: true},
		{name: "TestBookStorage_Delete", fields: fields{db: fakeDb}, args: args{id: 2}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := books.New(books.WithBookDb(tt.fields.db))
			if err := usr.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
