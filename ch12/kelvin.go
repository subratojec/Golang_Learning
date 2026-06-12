package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func main() {
	kelvin := 0.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin,"K is", celsius, "C")

	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Printf("%0.2f K is %0.2f F\n", kelvin, fahrenheit)

}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0/5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

// Do the functions.go experiments
