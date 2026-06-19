package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{
		{"Bradbury", -4.5895, 137.4417},
		{"Curiosity", -4.5895, 137.4523},
		{"Opportunity", -1.9462, -354.4734},
		{"Spirit", -14.6, -71.3},
	}

	for _, loc := range locations {
		fmt.Printf("%+v \n", loc)
	}
}
