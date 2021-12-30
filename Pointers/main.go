package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Greeen"

	log.Println("myString is set to", myString)
	changeMyStringWithPointer(&myString)
	log.Println("after function call myString is set to", myString)
}

func changeMyStringWithPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
