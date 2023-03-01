package car

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *carHandler) getCarsHandler(c echo.Context) error {
	cars, err := h.carUsecase.GetCars()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error_message": err.Error()})
	}
	return c.JSON(http.StatusOK, cars)
}
