package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var room = "lake"
	// switch {
	// case room == "cave":
	// 	fmt.Println("You find yourself in a daily lit cavern.")
	// case room == "lake":
	//  	fmt.Println("The ice seem solid enough.")
	// 	fallthrough
	// case room == "underwater":
	//  	fmt.Println("Water is freezing Cold.")
	// }
	//
	// shared code execution
	// day := "Tuesday"
	// switch day {
	// 	case "Monday","Tuesday","Wednesday","Thursday":
	// 	 	fmt.Println("Weekday")
	// 		fallthrough
	// 	case "Friday":
	// 	 	fmt.Println("Start of the Weekend")
	// 	case "Saturday", "Sunday":
	// 		fmt.Println("Weekend")
	// }


	// Short declaration in a switch statement
	switch num := rand.Intn(10); num{
		case 0:
			fmt.Println("Space Adventures")
		case 1:
			fmt.Println("Space X")
		case 2:
			fmt.Println("Virgin Galactic")
		default:
		 	fmt.Println("Random Spaceline #",num)
	}
}
