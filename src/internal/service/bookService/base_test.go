package bookService_test

import (
	"errors"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errStorageGet    = errors.New("storage get error")
	errStorageSet    = errors.New("storage set error")
	errStorageUpdate = errors.New("storage update error")
	errStorageDelete = errors.New("storage delete error")
	errStorageList   = errors.New("storage list error")
	errStorageFilter = errors.New("storage filter error")
)

type mockStorage struct {
	getErr    error
	setErr    error
	updateErr error
	deleteErr error
	listErr   error
	filterErr error
	bookDb    *mongo.Collection
}

func (m *mockStorage) Set(_ models.Book) (models.Book, error) {
	return models.Book{}, m.setErr
}

func (m *mockStorage) Get(_ int) (models.Book, error) {
	return models.Book{}, m.getErr
}

func (m *mockStorage) Update(_ int, _ map[string]interface{}) (models.Book, error) {
	return models.Book{}, m.updateErr
}

func (m *mockStorage) Delete(_ int) error {
	return m.deleteErr
}

func (m *mockStorage) List() ([]models.Book, error) {
	return nil, m.listErr
}

func (m *mockStorage) FilterBy(_ map[string]interface{}) ([]models.Book, error) {
	return nil, m.filterErr
}
