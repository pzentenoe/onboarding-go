package models

import "time"

type Car struct {
	ManufacturedYear int `json:"manufactured_year"`
}

func (c *Car) SetManufactureYear(year int) {
	c.ManufacturedYear = year
}

func (c *Car) GetYear() int {
	currentDate := time.Now().Year()
	return currentDate - c.ManufacturedYear
}

func (c *Car) GetYear2() int {
	currentDate := time.Now().Year()
	return currentDate - c.ManufacturedYear
}
