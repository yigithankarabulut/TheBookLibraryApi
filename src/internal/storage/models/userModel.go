package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	MemberNumber int                `json:"member_number" bson:"member_number"`
	FirstName    string             `json:"firstName" bson:"firstName"`
	LastName     string             `json:"lastName" bson:"lastName"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
}
