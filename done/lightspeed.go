package main

import "fmt"

func main() {

	const lightspeed = 299792 // km/s
	var distance = 56000000 // km

	fmt.Println(distance/lightspeed,"seconds")

	distance = 401000000
	fmt.Println(distance/lightspeed,"seconds")

	// Time for subtle to reach maras,
	distance = 96300000
	var subtle_speed = 100800
	fmt.Println("Time:",distance/subtle_speed,"h")
}
