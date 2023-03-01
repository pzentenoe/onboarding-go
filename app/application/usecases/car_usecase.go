package usecases

import "github.com/pzentenoe/onboarding-go/app/domain/models"

type CarUsecase interface {
	GetCars() ([]*models.Car, error)
	CreateCar(*models.Car) (*models.Car, error)
}
