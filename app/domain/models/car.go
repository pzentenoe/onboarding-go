package models

import "time"

type Car struct {
	ID               string `json:"id"`
	ManufacturedYear int    `json:"manufactured_year"`
	Description      string `json:"description"`
	Color            string `json:"color"`
	Plate            string `json:"plate"`
}

func (c *Car) SetManufactureYear(year int) {
	c.ManufacturedYear = year
}

func (c *Car) GetYear() int {
	currentDate := time.Now().Year()
	return currentDate - c.ManufacturedYear
}
