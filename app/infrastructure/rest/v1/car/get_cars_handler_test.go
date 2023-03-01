package car

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	getCarJSON      = "[{\"id\":\"1\",\"manufactured_year\":2000,\"description\":\"Mazda\",\"color\":\"Rojo\",\"plate\":\"HHJJ12\"}]\n"
	getCarJSONError = "{\"error_message\":\"error al obtener autos\"}\n"
)

func Test_carHandler_getCarsHandler(t *testing.T) {

	t.Run("when get cars works OK then returns car list", func(t *testing.T) {
		// Arrange
		e := echo.New()
		httpRequest := httptest.NewRequest(http.MethodGet, "/cars", nil)
		httpRequest.Header.Add("", "")
		resp := httptest.NewRecorder()
		echoContext := e.NewContext(httpRequest, resp)
		g := e.Group("/v1")

		mockCarUsecase := new(carUsecaseMock)
		mockCarUsecase.On("GetCars").Return(getCarsStub(), nil)

		h := NewCarHandler(mockCarUsecase, g)

		//Act
		err := h.getCarsHandler(echoContext)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, getCarJSON, resp.Body.String())
	})

	t.Run("when get cars works OK then returns car list", func(t *testing.T) {
		// Arrange
		e := echo.New()
		httpRequest := httptest.NewRequest(http.MethodGet, "/cars", nil)
		resp := httptest.NewRecorder()
		echoContext := e.NewContext(httpRequest, resp)
		g := e.Group("/v1")

		mockCarUsecase := new(carUsecaseMock)
		mockCarUsecase.On("GetCars").Return(nil, errors.New("error al obtener autos"))

		h := NewCarHandler(mockCarUsecase, g)

		//Act
		err := h.getCarsHandler(echoContext)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadGateway, resp.Code)
		assert.Equal(t, getCarJSONError, resp.Body.String())
	})

}
