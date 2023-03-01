package car

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pzentenoe/onboarding-go/app/domain/models"
	"github.com/pzentenoe/onboarding-go/app/infrastructure/database/mongo/car/entity"
	"github.com/pzentenoe/onboarding-go/app/infrastructure/database/mongo/car/mapper"
	logger "github.com/pzentenoe/onboarding-go/app/shared/log"
)

func (r *carMongoRepository) FindCars() ([]*models.Car, error) {
	log := logger.GetLogger()
	ctx := context.Background()
	collection := r.getCarCollection()
	cursor, err := collection.Find(ctx, primitive.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	carModels := make([]*models.Car, 0)
	for cursor.Next(ctx) {
		var car *entity.CarMongoEntity
		//car := &entity.CarMongoEntity{}
		if err = cursor.Decode(car); err != nil {
			log.WithError(err).Errorln()
			break
		}
		carModels = append(carModels, mapper.CarEntityToCarModel(car))
	}

	return carModels, nil
}
