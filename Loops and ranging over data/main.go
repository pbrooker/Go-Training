package main

import "log"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"dog", "fish","horse","cow","cat","bird"}
	// animals :=make(map[string]string)

	// animals["dog"] = "Buddy"
	// animals["cat"] = "Mezza"

	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// var sentence = "Once upon a midnight dreary"

	// for i, l := range sentence {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User

	users = append(users, User{"John", "Smith", "john@email.com", 30})
	users = append(users, User{"Alex", "Anderson", "alex@email.com", 20})
	users = append(users, User{"Frank", "Jones", "frank@email.com", 32})
	users = append(users, User{"Angie", "Martin", "angie@email.com", 40})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}