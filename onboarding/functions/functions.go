package main

import (
	"fmt"
	"github.com/pzentenoe/onboarding-go/onboarding/functions/interfaces"
	"github.com/pzentenoe/onboarding-go/onboarding/functions/models"
)

//Named return values
func main() {
	/*page, size, err := utils.ConvertPaginationValues("1", "10")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(page, size)*/

	car := &models.Car{ManufacturedYear: 2010}
	printYear(car)

	car.SetManufactureYear(2019)
	printYear(car)

	plane := &models.Plane{ManufacturedYear: 1999}
	printYear(plane)

}

func sumar(num1, num2 int) int {
	return num1 + num2
}

func printYear(vehicle interfaces.Vehicule) {
	years := vehicle.GetYear()
	fmt.Println(years)
}
