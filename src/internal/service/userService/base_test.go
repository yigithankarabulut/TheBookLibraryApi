package userService_test

import (
	"errors"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	errStorageSet    = errors.New("storage set error")
	errStorageGet    = errors.New("storage get error")
	errStorageUpdate = errors.New("storage update error")
	errStorageDelete = errors.New("storage delete error")
)

type mockStorage struct {
	setErr    error
	getErr    error
	updateErr error
	deleteErr error
	userDb    *mongo.Collection
}

func (m *mockStorage) Set(_ models.User) (models.User, error) {
	return models.User{}, m.setErr
}

func (m *mockStorage) Get(_ int) (models.User, error) {
	return models.User{}, m.getErr
}

func (m *mockStorage) Update(_ int, _ map[string]interface{}) (models.User, error) {
	return models.User{}, m.updateErr
}

func (m *mockStorage) Delete(_ int) error {
	return m.deleteErr
}

func (m *mockStorage) List() []models.User {
	return nil
}
