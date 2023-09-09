package users_test

import (
	"github.com/yigithankarabulut/libraryapi/database"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/users"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestUserStorage_Set(t *testing.T) {
	var fakeDb = database.FakeConnectUser()
	type fields struct {
		db *mongo.Collection
	}
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{name: "TestUserStorage_Set", fields: fields{db: fakeDb}, args: args{user: models.User{}}, want: models.User{}, wantErr: true},
		{name: "TestUserStorage_Set", fields: fields{db: fakeDb}, args: args{user: models.User{MemberNumber: 1, FirstName: "Test", LastName: "Test", Email: ""}}, want: models.User{MemberNumber: 1, FirstName: "Test", LastName: "Test", Email: ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usr := users.New(users.WithUserDb(tt.fields.db))
			got, err := usr.Set(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.MemberNumber != tt.want.MemberNumber {
				t.Errorf("Set() got = %v, want %v", got, tt.want)
			}
		})
	}
}
