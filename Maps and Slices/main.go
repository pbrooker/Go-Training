package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}


func (m *User) printName() (string, string) {
	return m.FirstName, m.LastName
}

func main() {

	// Maps
	myMap := make(map[string]string)

	myMap["dog"] = "Buddy"
	myMap["other-dog"] = "Paige"

	log.Println(myMap["dog"], myMap["other-dog"])

	mapTwo := make(map[string]int)

	mapTwo["first"] = 1
	mapTwo["second"] = 2

	log.Println(mapTwo["first"], mapTwo["second"])

	structMap := make(map[string]User)
	person := User{
		FirstName: "Joe",
		LastName: "Blow",
	}

	structMap["user1"] = person

	log.Println(structMap["user1"].FirstName, structMap["user1"].LastName)

	// Slices

	var mySlice []string
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	mySlice = append(mySlice, "Joe")
	mySlice = append(mySlice, "Rick")
	mySlice = append(mySlice, "Roll")

	var mySlice2 []int

	mySlice2 = append(mySlice2, 4)
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 3)
	sort.Ints(mySlice2)

	log.Println(mySlice, mySlice2, numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[5:8])
}