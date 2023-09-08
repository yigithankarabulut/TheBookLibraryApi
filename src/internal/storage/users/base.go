package users

import (
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	Set(user models.User) (models.User, error)
	Get(memberNumber int) (models.User, error)
	Update(memberNumber int, updateData map[string]interface{}) (models.User, error)
	Delete(memberNumber int) error
	List() []models.User
}

type userStorage struct {
	db *mongo.Collection
}

type UserStorageOption func(storage *userStorage)

func WithUserDb(user *mongo.Collection) UserStorageOption {
	return func(s *userStorage) {
		s.db = user
	}
}

func New(options ...UserStorageOption) UserStorer {
	ms := &userStorage{}

	for _, o := range options {
		o(ms)
	}
	return ms
}
