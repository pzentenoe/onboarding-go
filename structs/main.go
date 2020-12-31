package main

import (
	"encoding/json"
	"fmt"
	"github.com/pzentenoe/onboarding/structs/models"
)

type DiaSemana int

func main() {

	//Custom Types
	var lunes DiaSemana = 1
	fmt.Println(lunes)
	numero := 3
	total := int(lunes) + numero
	fmt.Println(total)

	//Structs
	daniela := &models.Persona{
		Nombre:   "Daniela",
		Apellido: "Mallea",
		Edad:     33,
		DatosPersonales: &models.DatosPersonales{
			Direccion: "assaassa",
			Telefono:  121221,
		},
	}

	fmt.Println(daniela.Nombre)
	fmt.Println(daniela.Edad)

	fmt.Println(daniela)

	fmt.Println("Es mayor de edad ?", daniela.MayorDeEdad())

	user1 := models.Usuario{
		UserName: "user1",
		Password: "12345",
		Persona:  daniela,
	}

	fmt.Println(user1.Persona.Edad)

	user2 := new(models.Usuario)
	pablo := new(models.Persona)
	datosPersonales := new(models.DatosPersonales)
	pablo.Nombre = "Pablo"
	pablo.DatosPersonales = datosPersonales
	user2.Persona = pablo

	//Conversion JSON

	danielaBytes, err := json.Marshal(daniela)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(danielaBytes))
	}

	//Desconversion
	danielaJSON := "{\"nombre\":\"Daniela\",\"apellido\":\"Mallea\",\"edad\":33,\"datos_personales\":{\"direccion\":\"assaassa\",\"telefono\":121221}}"

	var daniela2 *models.Persona
	err = json.Unmarshal([]byte(danielaJSON), &daniela2)
	if err != nil {
		fmt.Println("Error",err.Error())
	} else {
		fmt.Println(daniela2.Nombre)
		fmt.Println(daniela2.Edad)
		fmt.Println(daniela2.DatosPersonales.Direccion)
	}

}
