package main

import "fmt"

/*
 * When shouldn't three index sliciing be used ? unless you specifically
 * want to over-write the elements of the underlying array, it's far safer
 * to set the capacity with a three-index slice.
 */

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] // Length: 4, Cap:4
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
	terrestrial = planets[0:4] // Length: 4, Cap: 8
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
}
