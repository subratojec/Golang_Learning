package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// Function Declaration.
func kelvinToCelsius (k kelvin) celsius {
	return celsius(k - 273.15)
}

// Method Declaration
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//celsius converts F to C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5 / 9)
}

// Main function
func main() {
	var k kelvin = 294.0
	var c celsius
	c = kelvinToCelsius(k) // calls the kelvinToCelsius function
	c = k.celsius() // calls the celsius method
	fmt.Println(c)
}
