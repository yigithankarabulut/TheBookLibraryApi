# TheBookLibraryApi

## Project Files Tree
``` bash
.
├── cmd
│   └── server
│       └── main.go
├── database
│   ├── connectBookTest.go
│   ├── connectUserTest.go
│   └── dbconnect.go
├── go.mod
├── go.sum
├── LICENSE
└── src
    ├── apiserver
    │   └── apiserver.go
    └── internal
        ├── service
        │   ├── bookService
        │   │   ├── base.go
        │   │   ├── base_test.go
        │   │   ├── delete.go
        │   │   ├── delete_test.go
        │   │   ├── filterby.go
        │   │   ├── filterby_test.go
        │   │   ├── get.go
        │   │   ├── get_test.go
        │   │   ├── list.go
        │   │   ├── list_test.go
        │   │   ├── request.go
        │   │   ├── response.go
        │   │   ├── set.go
        │   │   ├── set_test.go
        │   │   ├── update.go
        │   │   └── update_test.go
        │   └── userService
        │       ├── base.go
        │       ├── base_test.go
        │       ├── delete.go
        │       ├── delete_test.go
        │       ├── get.go
        │       ├── get_test.go
        │       ├── list.go
        │       ├── list_test.go
        │       ├── request.go
        │       ├── response.go
        │       ├── set.go
        │       ├── set_test.go
        │       ├── update.go
        │       └── update_test.go
        ├── storage
        │   ├── books
        │   │   ├── base.go
        │   │   ├── delete.go
        │   │   ├── delete_test.go
        │   │   ├── filterby.go
        │   │   ├── filterby_test.go
        │   │   ├── get.go
        │   │   ├── get_test.go
        │   │   ├── list.go
        │   │   ├── list_test.go
        │   │   ├── set.go
        │   │   ├── set_test.go
        │   │   ├── update.go
        │   │   └── update_test.go
        │   ├── models
        │   │   ├── bookModel.go
        │   │   └── userModel.go
        │   └── users
        │       ├── base.go
        │       ├── delete.go
        │       ├── delete_test.go
        │       ├── get.go
        │       ├── get_test.go
        │       ├── list.go
        │       ├── list_test.go
        │       ├── set.go
        │       ├── set_test.go
        │       ├── update.go
        │       └── update_test.go
        └── transport
            └── http
                ├── basehttphandler
                │   └── basehttphandler.go
                ├── httpservice
                │   ├── base.go
                │   ├── base_test.go
                │   ├── createBook.go
                │   ├── createBook_test.go
                │   ├── deleteBook.go
                │   ├── deleteBook_test.go
                │   ├── filterBooks.go
                │   ├── filterBooks_test.go
                │   ├── getBook.go
                │   ├── getBooks.go
                │   ├── getBooks_test.go
                │   ├── getBook_test.go
                │   ├── loginUser.go
                │   ├── loginUser_test.go
                │   ├── logoutUser.go
                │   ├── logoutUser_test.go
                │   ├── registerUser.go
                │   ├── registerUser_test.go
                │   ├── router.go
                │   ├── updateBook.go
                │   └── updateBook_test.go
                └── middlewares
                    └── jwtmiddleware.go

```
