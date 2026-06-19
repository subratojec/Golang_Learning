package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
		// alt  float64
	}

	// var spirit location
	// spirit.lat = -14.6
	// spirit.long = -71.3
	// spirit.alt = -4400

	// spirit := location{lat: -14.6, long: -71.3}

	spirit := location{-14.5684, 175472636}

	// var opportunity location
	// opportunity.lat = -1.9462
	// opportunity.long = -354.4734
	// opportunity.alt = -4400

	// opportunity := location{lat: -1.9462, long: -354.4734}

	opportunity := location{-1.9462, -354.4734}

	// fmt.Println(spirit, opportunity)

	fmt.Printf("%v \n", spirit)
	fmt.Printf("%+v \n", opportunity)

	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.0106

	fmt.Println(bradbury, curiosity)
}
