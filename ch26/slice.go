package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto", "Haumea", "Makemake", "Eris",
	}
	reclassify(&planets)
	fmt.Println(planets)
}
