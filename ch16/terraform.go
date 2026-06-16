package main

import "fmt"

// terraform accomplishes nothing
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}
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
	dwarfs := [5]string{"Pluto", "Charon", "Eris", "Haumea", "Makemake"}
	terraform(planets)
	terraform(dwarfs)
	fmt.Println(planets)
	fmt.Println(dwarfs)
}
