package repositories

import "github.com/pzentenoe/onboarding-go/app/domain/models"

type CarRepository interface {
	FindCars() ([]*models.Car, error)
	CreateCar(*models.Car) (*models.Car, error)
}
