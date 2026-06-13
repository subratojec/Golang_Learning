package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// From celsius to two other temperature scales
func (c celsius) ToKelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) ToFahrenheit() fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

// From Kelvin to two other temp scales
func (k kelvin) ToCelsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) ToFahrenheit() fahrenheit {
	return fahrenheit(k*9/5 - 459.67)
}

// From Fahrenheit to two other temp scales
func (f fahrenheit) ToCelsius() celsius {
	return celsius((f - 32) * 5 / 9)
}
func (f fahrenheit) ToKelvin() kelvin {
	return kelvin((f + 459.67) * 5 / 9)
}

// Main function to demonstrate the temperature conversion methods
func main() {
	var c celsius = 0
	var k kelvin = 0
	var f fahrenheit = 0

	fmt.Println(c.ToKelvin(), c.ToFahrenheit())
	fmt.Println(k.ToCelsius(), k.ToFahrenheit())
	fmt.Println(f.ToCelsius(), f.ToKelvin())
}
