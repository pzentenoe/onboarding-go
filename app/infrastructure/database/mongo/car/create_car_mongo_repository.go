package car

import (
	"context"
	"github.com/pzentenoe/onboarding-go/app/domain/models"
	"github.com/pzentenoe/onboarding-go/app/infrastructure/database/mongo/car/mapper"
	logger "github.com/pzentenoe/onboarding-go/app/shared/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *carMongoRepository) CreateCar(carModel *models.Car) (*models.Car, error) {
	log := logger.GetLogger()
	carCollection := r.getCarCollection()
	carEntity := mapper.CarModelToCarEntity(carModel)

	result, err := carCollection.InsertOne(context.Background(), carEntity)
	if err != nil {
		log.WithError(err).Errorln("fail to create car")
		return nil, err
	}
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		carModel.ID = objectID.Hex()
	}
	return carModel, nil
}
