package books

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (book *bookStorage) Update(id int, updateData map[string]interface{}) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := book.Get(id); err != nil {
		return models.Book{}, err
	}
	update := bson.M{
		"$set": updateData,
	}
	updateRes, err2 := book.db.UpdateOne(
		ctx,
		bson.M{"id": id},
		update,
	)
	if err2 != nil || updateRes.ModifiedCount == 0 {
		return models.Book{}, err2
	}
	returnBook, err3 := book.Get(id)
	if err3 != nil {
		return models.Book{}, err3
	}
	return returnBook, nil
}
