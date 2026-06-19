package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
		alt  float64
	}
	curiosity.lat = 37.7749
	curiosity.long = -122.4194
	curiosity.alt = -4400

	fmt.Println(curiosity)
}
