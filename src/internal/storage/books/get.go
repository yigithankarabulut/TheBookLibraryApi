package books

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (book *bookStorage) Get(id int) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var books models.Book
	if err := book.db.FindOne(ctx, bson.M{"id": id}).Decode(&books); err != nil {
		return models.Book{}, err
	}
	return books, nil
}
