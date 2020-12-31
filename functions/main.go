package main

import (
	"fmt"
	"github.com/pzentenoe/onboarding-go/structs/models"
	"strconv"
	"time"
)

func main() {

	p1 := &models.Persona{
		Nombre: "Daniela",
	}

	fmt.Println(p1.Nombre)
	cambiarNombre("Pablo", p1)

	fmt.Println(p1.Nombre)

	d, m, y := GetDayMonthYear(time.Now())
	fmt.Println(d)
	fmt.Println(m)
	fmt.Println(y)

}

func cambiarNombre(nombre string, persona *models.Persona) {
	persona.Nombre = nombre
}

func GetDayMonthYear(dataTime time.Time) (day int, monthNumber int, year int) {
	month := dataTime.Format("01")
	monthNumber, _ = strconv.Atoi(month)

	year, _, day = dataTime.Date()
	return
}

func GetUserById(id string) (*models.Usuario, error) {
	//llamar servicio
	return &models.Usuario{
		UserName: "User1",
		Password: "123456",
	}, nil
}
