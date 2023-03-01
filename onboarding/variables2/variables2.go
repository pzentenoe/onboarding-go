package main

import (
	"fmt"
	"strconv"
)

func main() {
	//palabra := "asjkasjljkasjaskla jl"

	var numeroString string
	numeroString = "100"

	if numero, err := strconv.Atoi(numeroString); err != nil {
		fmt.Println(numero)
		fmt.Println(err.Error())
	} else {
		fmt.Println(numero)
		fmt.Println(strconv.Itoa(numero))

	}

	//var runeArray = []rune(palabra)

	//For I
	/*	for i := 0; i < len(runeArray); i++ {
		fmt.Println(strconv.QuoteRune(runeArray[i]))
	}*/

	//For range - foreach
	/*	for i, letra := range runeArray {
		fmt.Println(i)
		fmt.Println(strconv.QuoteRune(letra))
	}*/

	for {
		fmt.Println("While true")
		break
	}

	for 3 != 4 {
		fmt.Println("Ejecutando")
		break
	}

}
