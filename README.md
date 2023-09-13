# TheBookLibraryApi
TheBookLibraryApi is a library management application

## Getting Started

In this section, provide information on how to get your project up and running in a local environment.

### Requirements

To run the project, you'll need the following:

- [Go](https://golang.org/) (at least version 1.21.0)
- [Fiber](https://github.com/gofiber/fiber) web framework
- [MongoDB](https://github.com/mongodb/mongo-go-driver) Nosql Databases
- [Jwt](https://github.com/golang-jwt/jwt) JSON Web Tokens.
  
### Installation

To start the project, follow these steps:

- Replace the connection key in the database/{dbconnect,connectBookTest,connectUserTest}.go file with your own mongodb connection key

-Navigate to the project directory:

   ```bash
   1. cd TheBookLibraryApi
   2. Install project dependencies by running: go get {fiber,mongo,jwt}
   3. Start the application: go run cmd/main.go
   ```
## Endpoints
```http
POST    /users
POST    /login
GET     /logout

POST    /books/filter            "{ Searches for books with a filter }"
POST    /books                   "{ Create a book }"
GET     /books                   "{ Get all the books }"
GET     /book/?id=123            "{ Get book by id }"
PUT     /book/?id=123            "{ Update book by id }"
DELETE  /book/?id=123            "{ Delete book by id }"

```


## Json Examples

POST /books
```json
{
    "id":171,
    "title":"PostmanPostBookRequest",
    "author":"Yigithan",
    "isbn":"112233445566",
    "pageCount":250,
    "category":"rest",
    "language":"golang"
}
```

POST /books/filter
```json
{
    "category": "golang",
    "pageCount": 250
}
```
UPDATE /book/?id=171
```json
{
    "title":"Github",
    "category":"embeded",
    "language":"C",
    "pageCount":750
}
```
POST /users
```json
{
    "member_number":12,
    "firstName":"Yigithan",
    "lastName":"Karabulut",
    "email":"yigithannkarabulutt@gmail.com",
    "password":"1234"
}
```
POST /login
```json
{
    "member_number":12,
    "password":"1234"
}
```
### Project Files Tree
```bash
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

18 directories, 87 files
```
