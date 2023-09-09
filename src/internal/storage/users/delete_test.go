package users_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserStorage_Delete(t *testing.T) {
	var fakeDb = database.FakeConnectUser()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		memberNumber int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestUserStorage_Delete", fields: fields{db: fakeDb}, args: args{memberNumber: 0}, wantErr: true},
		{name: "TestUserStorage_Delete", fields: fields{db: fakeDb}, args: args{memberNumber: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := users.New(users.WithUserDb(tt.fields.db))
			if err := usr.Delete(tt.args.memberNumber); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
