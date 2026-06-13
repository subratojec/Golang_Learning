package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}
func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	k2 := celsiusToKelvin(c)
	fmt.Printf("%0.2f\n", c)
	fmt.Printf("%0.2f\n", k2)
}
