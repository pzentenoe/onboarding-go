package main

import "fmt"

func main() {

	//Array
	var nombres [6]string
	nombres = [6]string{"Pablo", "Daniela", "Manuel", "Juan", "Ale", "Diego"}

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}
	fmt.Println("---------------------")
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
	fmt.Println("---------------------")
	i := 0

	for {
		fmt.Println(i + 1)
		if i == 10 {
			break
		}
		i++
	}

	j := 0

	fmt.Println("---------------------")

	for j < 10 {
		fmt.Println(j + 1)
		j++
	}

	//Slice
	fmt.Println("--------SLICES-------------")
	apellidos2 := make([]string, 0)
	apellidos2 = append(apellidos2, "Zenteno", "Mallea")

	apellidos := make([]string, 0)
	fmt.Println(len(apellidos))
	fmt.Println(cap(apellidos))

	apellidos = append(apellidos, apellidos2...)
	apellidos = append(apellidos, "Labarca")



	for _, apellido := range apellidos2 {
		fmt.Println(apellido)
	}
	fmt.Println("---------------------")
	for _, apellido := range apellidos {
		fmt.Println(apellido)
	}

}
