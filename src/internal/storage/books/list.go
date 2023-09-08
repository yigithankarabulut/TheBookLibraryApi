package books

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (book *bookStorage) List() ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := book.db.Find(ctx, bson.M{})

	var books []models.Book
	if err = result.All(ctx, &books); err != nil {
		return []models.Book{}, err
	}
	return books, nil
}
