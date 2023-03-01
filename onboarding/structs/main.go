package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/pzentenoe/onboarding-go/onboarding/structs/models"
)

const personJSON = `{"id":"asasasasas","name":"Felipe","last_name":"Montero","age":0,"is_active":false}`

func main() {

	/*	user1 := models.User{}
		user1.ID = "asasas"
		user1.Username = "pzenteno"

		fmt.Println(user1)*/

	person1 := models.Person{
		ID:       "asasasasas",
		Name:     "Felipe",
		LastName: "Montero",
	}

	jsonBytes, err := json.Marshal(person1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var person2 *models.Person

	json.Unmarshal([]byte(personJSON), &person2)
	fmt.Println(person2)

	xmlBytes, _ := xml.Marshal(person2)

	fmt.Println(string(jsonBytes))
	fmt.Println(string(xmlBytes))

}
