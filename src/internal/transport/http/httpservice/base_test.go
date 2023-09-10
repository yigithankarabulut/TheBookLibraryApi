package httpservice_test

import (
	"context"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/bookService"
	"github.com/yigithankarabulut/libraryapi/src/internal/service/userService"
	"github.com/yigithankarabulut/libraryapi/src/internal/storage/models"
)

type mockBookService struct {
	book               models.Book
	bookSetErr         error
	bookGetErr         error
	bookUpdateErr      error
	bookDeleteErr      error
	bookListErr        error
	bookFilterErr      error
	bookRequest        *bookService.SetBookRequest
	bookResponse       *bookService.BookResponse
	bookUpdateRequest  *bookService.UpdateBookRequest
	bookListResponse   *bookService.ListBooksResponse
	bookFilterResponse *bookService.ListBooksResponse
}

type mockUserService struct {
	user              models.User
	userSetErr        error
	userGetErr        error
	userUpdateErr     error
	userDeleteErr     error
	userListErr       error
	userRequest       *userService.SetUserRequest
	userResponse      *userService.UserResponse
	userUpdateRequest *userService.UpdateUserRequest
	userListResponse  *userService.ListUserResponse
}

// userStoreService methods

func (m *mockUserService) Set(_ context.Context, _ *userService.SetUserRequest) (*userService.UserResponse, error) {
	return m.userResponse, m.userSetErr
}

func (m *mockUserService) Get(_ context.Context, _ int) (*userService.UserResponse, error) {
	return m.userResponse, m.userGetErr
}

func (m *mockUserService) Update(_ context.Context, _ *userService.UpdateUserRequest) (*userService.UserResponse, error) {
	return m.userResponse, m.userUpdateErr
}

func (m *mockUserService) Delete(_ context.Context, _ int) error {
	return m.userDeleteErr
}

func (m *mockUserService) List(_ context.Context) (*userService.ListUserResponse, error) {
	return m.userListResponse, m.userListErr
}

// bookStoreService methods

func (m *mockBookService) Set(_ context.Context, _ *bookService.SetBookRequest) (*bookService.BookResponse, error) {
	return m.bookResponse, m.bookSetErr
}

func (m *mockBookService) Get(_ context.Context, _ int) (*bookService.BookResponse, error) {
	return m.bookResponse, m.bookGetErr
}

func (m *mockBookService) Update(_ context.Context, _ *bookService.UpdateBookRequest) (*bookService.BookResponse, error) {
	return m.bookResponse, m.bookUpdateErr
}

func (m *mockBookService) Delete(_ context.Context, _ int) error {
	return m.bookDeleteErr
}

func (m *mockBookService) List(_ context.Context) (*bookService.ListBooksResponse, error) {
	return m.bookListResponse, m.bookListErr
}

func (m *mockBookService) FilterBy(_ context.Context, _ *bookService.FilterByBookRequest) (*bookService.ListBooksResponse, error) {
	return m.bookFilterResponse, m.bookFilterErr
}
