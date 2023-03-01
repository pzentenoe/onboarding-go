package car

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type carMongoRepository struct {
	mongoClient *mongo.Client
}

func NewCarMongoRepository(mongoClient *mongo.Client) *carMongoRepository {
	return &carMongoRepository{mongoClient: mongoClient}
}

func (r *carMongoRepository) getCarCollection() *mongo.Collection {
	collection := r.mongoClient.Database("automotive").Collection("cars")
	return collection
}
