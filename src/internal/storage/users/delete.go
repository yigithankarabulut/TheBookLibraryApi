package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (usr *userStorage) Delete(memberNumber int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := usr.Get(memberNumber); err != nil {
		return err
	}
	res, err := usr.db.DeleteOne(ctx, bson.M{"member_number": memberNumber})
	if err != nil || res.DeletedCount == 0 {
		return err
	}
	return nil
}
