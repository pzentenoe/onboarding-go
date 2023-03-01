package models

import "time"

type Plane struct {
	ManufacturedYear int `json:"manufactured_year"`
}

func (c *Plane) SetManufactureYear(year int) {
	c.ManufacturedYear = year
}

func (c *Plane) GetYear() int {
	currentDate := time.Now().Year()
	return currentDate - c.ManufacturedYear
}

func (c *Plane) GetYear2() int {
	currentDate := time.Now().Year()
	return currentDate - c.ManufacturedYear
}
