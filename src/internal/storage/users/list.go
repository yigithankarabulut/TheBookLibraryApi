package users

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (usr *userStorage) List() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := usr.db.Find(ctx, bson.M{})

	var users []models.User
	if err = result.All(ctx, &users); err != nil {
		return []models.User{}
	}
	return users
}
