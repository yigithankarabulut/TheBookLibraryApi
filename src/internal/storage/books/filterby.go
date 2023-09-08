package books

import (
	"context"
	"fmt"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (book *bookStorage) FilterBy(filter map[string]interface{}) ([]models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var books []models.Book
	filters := bson.M{}
	if title, ok := filter["title"]; ok && title != nil {
		filters["title"] = title
	}
	if author, ok := filter["author"]; ok && author != nil {
		filters["author"] = author
	}
	if publishDate, ok := filter["publishDate"]; ok && publishDate != nil {
		filters["publishDate"] = publishDate
	}
	if isbn, ok := filter["isbn"]; ok && isbn != nil {
		filters["isbn"] = isbn
	}
	if pageCount, ok := filter["pageCount"]; ok && pageCount != nil {
		filters["pageCount"] = pageCount
	}
	if category, ok := filter["category"]; ok && category != nil {
		filters["category"] = category
	}
	if language, ok := filter["language"]; ok && language != nil {
		filters["language"] = language
	}
	if description, ok := filter["description"]; ok && description != nil {
		filters["description"] = description
	}
	result, err := book.db.Find(ctx, filters)
	if err != nil {
		return []models.Book{}, err
	}
	if err = result.All(ctx, &books); err != nil {
		return []models.Book{}, err
	}
	if len(books) == 0 {
		return []models.Book{}, fmt.Errorf("no books found")
	}
	return books, nil
}
