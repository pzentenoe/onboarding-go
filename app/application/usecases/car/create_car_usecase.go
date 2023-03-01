package car

import "github.com/pzentenoe/onboarding-go/app/domain/models"

func (u *carUsecase) CreateCar(carModel *models.Car) (*models.Car, error) {
	return u.carRepository.CreateCar(carModel)
}
