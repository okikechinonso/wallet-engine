package entity

import "time"

type Movie struct {
	ID    string    `json:"id" bson:"_id"`
	Name  string    `json:"name" bson:"name"`
	Email string    `json:"email" bson:"email"`
	Text  string    `json:"text" bson:"text"`
	Date  time.Time `json:"time" bson:"date"`
}
