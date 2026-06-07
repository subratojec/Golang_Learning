package main

import "fmt"

func main() {
	var choice = "cave"
	if choice == "cave" {
		fmt.Println("You find yourself in a dimly lit cave.")
	} else if choice == "mountain" {
		fmt.Println("You head further up the mountain.")
	} else {
		fmt.Println("You are at the Entrance.")
	}
}
