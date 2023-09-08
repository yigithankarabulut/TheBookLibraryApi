package books

import (
	"context"
	"fmt"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"time"
)

func (book *bookStorage) Set(books models.Book) (models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := book.Get(books.Id); err == nil {
		return models.Book{}, fmt.Errorf("book with id: %d already exists", books.Id)
	}
	insertResult, err := book.db.InsertOne(ctx, books)
	if err != nil || insertResult.InsertedID == nil {
		return models.Book{}, err
	}
	return books, nil
}
