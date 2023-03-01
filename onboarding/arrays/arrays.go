package main

import "fmt"

func main() {

	var letras [27]string
	letras = [27]string{}
	letras[0] = "a"
	letras[1] = "b"
	letras[2] = "b"

	var vocalesSlice = make([]string, 0)
	var letras2Slice = make([]string, 0)
	vocalesSlice = append(vocalesSlice, "a", "e", "i", "o", "u")
	letras2Slice = append(letras2Slice, vocalesSlice[:3]...)
	//letras2Slice = append(letras2Slice, vocalesSlice[3:]...)

	fmt.Println(vocalesSlice[:3])
	fmt.Println("letras2Slice", letras2Slice)
	fmt.Println(len(vocalesSlice))
	fmt.Println(cap(vocalesSlice))

	numeros := []int{}
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	numeros = append(numeros, 6, 7, 8)

}
