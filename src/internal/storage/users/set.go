package users

import (
	"context"
	"fmt"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"time"
)

func (usr *userStorage) Set(user models.User) (models.User, error) {
	if _, err := usr.Get(user.MemberNumber); err == nil {
		return models.User{}, fmt.Errorf("err: user already exists")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if user == (models.User{}) {
		return models.User{}, fmt.Errorf("err: user is empty")
	}
	insertResult, err := usr.db.InsertOne(ctx, user)
	if err != nil || insertResult.InsertedID == nil {
		return models.User{}, fmt.Errorf("err: user could not be inserted")
	}
	return user, nil
}
