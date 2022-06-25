package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID    string             `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
	Text  string             `json:"text" bson:"text"`
	Date  primitive.DateTime `json:"time" bson:"date"`
}
