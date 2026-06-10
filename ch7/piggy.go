package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nickels := 5
	dimes := 10
	quarters := 25

	amount := 0
	for amount <= 2000 {
		num := rand.Intn(3) + 1
		switch num {
			case 1:
			 	amount += nickels
			case 2:
			 	amount += dimes
			case 3:
			 	amount += quarters
		}
		fmt.Printf("$%d,%02d\n", amount/100, amount%100)
  	}
}
