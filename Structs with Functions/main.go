package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Joe"

	myVar2 := myStruct{
		FirstName: "Nick",
	}

	log.Println("myVar set to", myVar.printFirstName())
	log.Println("myVar2 set to", myVar2.printFirstName())

}