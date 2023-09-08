package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	UID         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Id          int                `json:"id" bson:"id"`
	Title       string             `json:"title" bson:"title"`
	Author      string             `json:"author" bson:"author"`
	PublishDate time.Time          `json:"publishDate,omitempty" bson:"publishDate,omitempty"`
	ISBN        string             `json:"isbn" bson:"isbn"`
	PageCount   int                `json:"pageCount" bson:"pageCount"`
	Category    string             `json:"category" bson:"category"`
	Language    string             `json:"language" bson:"language"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
