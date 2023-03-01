package car

import (
	"github.com/pzentenoe/onboarding-go/app/domain/repositories"
)

type carUsecase struct {
	carRepository repositories.CarRepository
}

func NewCarUsecase(carRepository repositories.CarRepository) *carUsecase {
	return &carUsecase{carRepository: carRepository}
}
