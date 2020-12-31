package main

import "fmt"

func main() {
	//maps
	nombresEdadMap := make(map[string]int)
	nombresEdadMap["Pablo"] = 32
	nombresEdadMap["Manuel"] = 25
	nombresEdadMap["Daniela"] = 33

	if edad, ok := nombresEdadMap["Peter"]; ok {
		fmt.Printf("Fue encontrada la edad de Daniela %v", edad)
	} else {
		fmt.Println("No fue encontrado")
	}

	for key, value := range nombresEdadMap {
		fmt.Println("Key ", key)
		fmt.Println("Value ", value)
	}

}
