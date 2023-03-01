package main

import "fmt"

func main() {
	numerosInt := make(map[int]string)
	numerosInt[1] = "uno"
	numerosInt[2] = "dos"
	numerosInt[3] = "tres"

	numerosString := make(map[string]int)
	numerosString["uno"] = 1
	numerosString["dos"] = 2
	numerosString["tres"] = 3

	if value, ok := numerosString["tres"]; ok {
		fmt.Println(value)
	}

	for _, v := range numerosString {
		fmt.Println(v)
		fmt.Println(numerosInt[v])
	}

	delete(numerosInt, 1)

	fmt.Println(numerosInt)

	paises := map[string]string{"AR": "Argentina"}
	paises["CL"] = "Chile"

}
