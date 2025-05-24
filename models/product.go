package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Price float64            `bson:"price" json:"price"`
}
