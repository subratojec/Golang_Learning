package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	planetMark2 := planets
	planets[2] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetMark2)
}
