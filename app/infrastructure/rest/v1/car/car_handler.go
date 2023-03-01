package car

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-go/app/application/usecases"
)

type carHandler struct {
	carUsecase usecases.CarUsecase
}

func NewCarHandler(carUsecase usecases.CarUsecase, e *echo.Group) *carHandler {
	h := &carHandler{carUsecase: carUsecase}
	e.GET("/cars", h.getCarsHandler)
	e.POST("/cars", h.createCarHandler)
	return h
}
