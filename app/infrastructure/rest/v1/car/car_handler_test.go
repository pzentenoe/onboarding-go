package car

import (
	"github.com/pzentenoe/onboarding-go/app/domain/models"
	"github.com/stretchr/testify/mock"
)

type carUsecaseMock struct {
	mock.Mock
}

func (m *carUsecaseMock) GetCars() ([]*models.Car, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Car), args.Error(1)
}

func (m *carUsecaseMock) CreateCar(car *models.Car) (*models.Car, error) {
	args := m.Called(car)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Car), args.Error(1)
}

func getCarsStub() []*models.Car {
	cars := make([]*models.Car, 0)
	cars = append(cars, getCarStub())
	return cars
}

func getCarStub() *models.Car {
	return &models.Car{
		ID:               "1",
		ManufacturedYear: 2000,
		Description:      "Mazda",
		Color:            "Rojo",
		Plate:            "HHJJ12",
	}
}
