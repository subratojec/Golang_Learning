package main

import "fmt"

func main() {
	var administrator *string
	scolese := "J. Scolese"

	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)

	fmt.Println(administrator == major)

	// After the clone is made, direct and indirect modifications to
	// bolden have no effect on the value of charles, or vice versa.
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)
}
