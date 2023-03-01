package car

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-go/app/domain/models"
	"github.com/pzentenoe/onboarding-go/app/shared/log"
	"net/http"
)

func (h *carHandler) createCarHandler(c echo.Context) error {
	log := log.GetLogger()
	var carModel *models.Car
	if err := c.Bind(&carModel); err != nil {
		log.WithError(err).Println("invalid body")
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid body"})
	}
	newCar, err := h.carUsecase.CreateCar(carModel)
	if err != nil {
		log.WithError(err).Println(err.Error())
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, newCar)
}
