package users

import (
	"context"
	"fmt"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (usr *userStorage) Get(memberNumber int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	if err := usr.db.FindOne(ctx, bson.M{"member_number": memberNumber}).Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("err: user not found")
	}
	return user, nil
}
