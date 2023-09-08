package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (book *bookStorage) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := book.Get(id); err != nil {
		return err
	}
	res, err := book.db.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || res.DeletedCount == 0 {
		return err
	}
	return nil
}
