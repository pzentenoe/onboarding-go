package mapper

import (
	"github.com/pzentenoe/onboarding-go/app/domain/models"
	"github.com/pzentenoe/onboarding-go/app/infrastructure/database/mongo/car/entity"
)

func CarEntityToCarModel(car *entity.CarMongoEntity) *models.Car {
	return &models.Car{
		ID:               car.ID.Hex(),
		ManufacturedYear: car.Year,
		Description:      car.Description,
		Color:            car.Color,
		Plate:            car.Plate,
	}
}

func CarModelToCarEntity(car *models.Car) *entity.CarMongoEntity {
	return &entity.CarMongoEntity{
		Year:        car.ManufacturedYear,
		Description: car.Description,
		Color:       car.Color,
		Plate:       car.Plate,
	}
}
