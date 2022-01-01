package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Colour        string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Buddy",
		Breed: "Pyrenees",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Bob",
		Colour: "black",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "bark!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "grunt!"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}