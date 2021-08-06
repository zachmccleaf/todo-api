package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct

type Todo struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn       string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Message    string             `json:"message" bson:"message,omitempty"`
	Color      string             `json:"color" bson:"color,omitempty"`
	IsComplete bool               `json:"isComplete" bson:"isComplete,omitempty"`
}
