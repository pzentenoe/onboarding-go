package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Imprimir en consola
	fmt.Println("Hola Mundo")

	//Variables
	var numeroPositivo uint8
	numeroPositivo = 10
	fmt.Println(numeroPositivo)

	numeroInt := 10
	fmt.Println(numeroInt)

	var numeroFloat32 float32
	numeroFloat32 = 32.5
	fmt.Println(numeroFloat32)

	numeroFloat64 := 32.5
	fmt.Println(numeroFloat64)

	isTrue := true
	fmt.Println(isTrue)

	var isAlgo bool
	fmt.Println(isAlgo)

	character := 'a'

	var letra rune
	letra = 'a'
	fmt.Println(strconv.QuoteRune(letra))

	fmt.Println(character)

	r := []rune("ABC€")
	fmt.Println(r)
	fmt.Printf("%U\n", r)
	fmt.Println(string(r))
}
