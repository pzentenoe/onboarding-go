package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CarMongoEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description"`
	Year        int                `bson:"year"`
	Color       string             `bson:"color"`
	Plate       string             `bson:"plate"`
}
