package main

import "fmt"

func main() {
	// fmt.Println("There is a cavern entrance here and a path to the east.")
	// var command = "go inside"
	// switch command {
	// 	case "go east":
	// 		fmt.Println("You head further up the mountain.")
	// 	case "enter cave", "go inside":
	// 		fmt.Println("You find yourself in a dimly lit cavern.")
	// 	case "read sign":
	// 		fmt.Println("The sign reads 'No Minors'.")
	// 	default:
	// 		fmt.Println("Didn't quite get that.")
	// }


	// Switch can use the number.
	day := 0
	switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		default:
			fmt.Println("Invalid day")
	}
}
