package users_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserStorage_List(t *testing.T) {
	var fakeDb = database.FakeConnectUser()
	type fields struct {
		db *mongo.Collection
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "TestUserStorage_List", fields: fields{db: fakeDb}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := users.New(users.WithUserDb(tt.fields.db))
			if got := usr.List(); len(got) < 1 {
				t.Errorf("List() = %v", got)
			}
		})
	}
}
