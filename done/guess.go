package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 20
	var guessed_number = rand.Intn(100)
	if guessed_number <= number {
  		fmt.Println("Number is too small")
	} else if guessed_number >= number {
		fmt.Println("Number is too big")
	} else {
		fmt.Println("Number is just right")
	}
}
