package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	HairColour string `json:"hair_colour"`
	HasDog     bool   `json:"has_dog"`
}

func main() {
	myJson := `[
		{
			"first_name": "Clark",
			"last_name": "Griswald",
			"hair_colour": "black",
			"has_dog": false
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_colour": "black",
			"has_dog": true
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Clark"
	m1.LastName = "Kent"
	m1.HairColour = "black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColour = "black"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}