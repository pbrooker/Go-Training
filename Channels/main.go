package main

import (
	"log"

	"github.com/pbrooker/gotrainingpackage/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randonNumber := helpers.RandomNumber(numPool)
	intChan <- randonNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <- intChan

	log.Println(num)
}
