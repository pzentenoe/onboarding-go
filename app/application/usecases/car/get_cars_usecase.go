package car

import "github.com/pzentenoe/onboarding-go/app/domain/models"

func (u *carUsecase) GetCars() ([]*models.Car, error) {
	cars, err := u.carRepository.FindCars()
	if err != nil {
		return nil, err
	}
	return cars, nil
}
