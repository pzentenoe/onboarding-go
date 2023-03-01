package car

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-go/app/application/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CarTableTest struct {
	description string
	usecase     usecases.CarUsecase
	c           echo.Context
	assert      assert.ErrorAssertionFunc
	comparison  assert.ComparisonAssertionFunc
}

func Test_carHandler_createCarHandler(t *testing.T) {

	car := getCarStub()
	carBytes, _ := json.Marshal(car)

	// Arrange
	e := echo.New()
	httpRequest := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewBuffer(carBytes))
	httpRequest.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	echoContext := e.NewContext(httpRequest, resp)
	g := e.Group("/v1")

	mockCarUsecaseOK := createCarOK()

	tests := []CarTableTest{
		{
			description: "when create car works OK",
			usecase:     mockCarUsecaseOK,
			c:           echoContext,
			assert:      assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			h := NewCarHandler(tt.usecase, g)
			err := h.createCarHandler(tt.c)
			tt.assert(t, err)
		})
	}
}

func createCarOK() *carUsecaseMock {
	mockCarUsecaseOK := new(carUsecaseMock)
	mockCarUsecaseOK.On("CreateCar", mock.AnythingOfType("*models.Car")).
		Return(getCarStub(), nil)
	return mockCarUsecaseOK
}
func createCarError() *carUsecaseMock {
	mockCarUsecaseError := new(carUsecaseMock)
	mockCarUsecaseError.On("CreateCar", mock.AnythingOfType("*models.Car")).
		Return(nil, errors.New("Error"))
	return mockCarUsecaseError
}
