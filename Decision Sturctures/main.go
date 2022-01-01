package main

import "log"

func main() {

	isTrue := true

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTre is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is false")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 100 and isTrue is true")
	} else if myNum > 100 || isTrue {
		log.Println("myNum is greater than 100b or isTrue is true")
	}

	myVar := "cat"

	switch myVar {
	case "dog":
		log.Println("myVar is set to dog")
	case "owl":
		log.Println("myVar is set to owl")
	case "fish":
		log.Println("myVar is set to fish")
	case "cat":
		log.Println("myVar is set to cat")
	default:
		log.Println("myVar is set to", myVar)
	}
}