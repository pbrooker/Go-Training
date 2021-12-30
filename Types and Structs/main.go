package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

func main() {
	// var s2 = "six"

	// log.Println("s is", s)
	// log.Println("s2 is", s2)

	// saySomthing("xxx")

	user := User {
		FirstName: "Joe",
		LastName: "Blow",
		PhoneNumber: "555-111-2222",
		Age: 50,
	}

	log.Println(user.FirstName, user.LastName, "Birthdate", user.Birthdate)
}
