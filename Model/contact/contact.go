package contact

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	PhoneNumber int `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}
