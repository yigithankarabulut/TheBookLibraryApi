package books

import (
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookStorer interface {
	Set(book models.Book) (models.Book, error)
	Get(id int) (models.Book, error)
	Update(id int, updateData map[string]interface{}) (models.Book, error)
	Delete(id int) error
	List() ([]models.Book, error)
	FilterBy(filter map[string]interface{}) ([]models.Book, error)
}

type bookStorage struct {
	db *mongo.Collection
}

type BookStorageOption func(storage *bookStorage)

func WithBookDb(book *mongo.Collection) BookStorageOption {
	return func(s *bookStorage) {
		s.db = book
	}
}

func New(options ...BookStorageOption) BookStorer {
	ms := &bookStorage{}

	for _, o := range options {
		o(ms)
	}
	return ms
}
