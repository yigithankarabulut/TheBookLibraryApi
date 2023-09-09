package users_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserStorage_Update(t *testing.T) {
	var fakeDb = database.FakeConnectUser()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		memberNumber int
		updateData   map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestUserStorage_Update", fields: fields{db: fakeDb}, args: args{memberNumber: 0, updateData: map[string]interface{}{"first_name": "Test", "last_name": "Test", "email": ""}}, wantErr: true},
		{name: "TestUserStorage_Update", fields: fields{db: fakeDb}, args: args{memberNumber: 2, updateData: map[string]interface{}{"first_name": "Test", "last_name": "Test", "email": ""}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := users.New(users.WithUserDb(tt.fields.db))
			if _, err := usr.Update(tt.args.memberNumber, tt.args.updateData); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
