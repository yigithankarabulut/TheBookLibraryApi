package books_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/books"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestBookStorage_Update(t *testing.T) {
	var fakeDb = database.FakeConnectBook()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		id         int
		updateData map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestBookStorage_Update", fields: fields{db: fakeDb}, args: args{id: 0, updateData: map[string]interface{}{"title": "Test", "author": "Test", "isbn": ""}}, wantErr: true},
		{name: "TestBookStorage_Update", fields: fields{db: fakeDb}, args: args{id: 1, updateData: map[string]interface{}{"title": "Test", "author": "Test", "isbn": "Test"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := books.New(books.WithBookDb(tt.fields.db))
			if _, err := usr.Update(tt.args.id, tt.args.updateData); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
