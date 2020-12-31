package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//uint
	var numeroUint uint = 1000
	fmt.Println(numeroUint)

	//int
	var numeroInt int
	numeroInt = 1000
	fmt.Println(numeroInt)

	//float
	numeroFloat := 200.2
	fmt.Println(numeroFloat)

	var numeroFloat32 float32 = 200.3
	fmt.Println(numeroFloat32)

	//bool
	isJueves := true
	fmt.Println(isJueves)

	//rune
	var letra rune
	letra = 's'
	fmt.Println(letra)

	//String
	palabra := "aasashkash,jklashakjsh"

	palabras := strings.Split(palabra, ",")

	fmt.Println(palabras[0])
	fmt.Println(palabras[1])

	numeroString := "10assas"
	numero, err := strconv.Atoi(numeroString)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(numero)





}
