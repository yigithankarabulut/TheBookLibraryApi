package users

import (
	"context"
	"fmt"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (usr *userStorage) Update(memberNumber int, updateData map[string]interface{}) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := usr.Get(memberNumber); err != nil {
		return models.User{}, fmt.Errorf("err: user not found")
	}
	update := bson.M{
		"$set": updateData,
	}
	updateRes, err2 := usr.db.UpdateOne(
		ctx,
		bson.M{"member_number": memberNumber},
		update,
	)
	if err2 != nil || updateRes.ModifiedCount == 0 {
		return models.User{}, err2
	}
	returnUser, err3 := usr.Get(memberNumber)
	if err3 != nil {
		return models.User{}, err3
	}
	return returnUser, nil
}
