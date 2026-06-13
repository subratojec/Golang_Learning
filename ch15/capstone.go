package main

import "fmt"

type fahrenheit float64
type celsius float64

// conversion function
func celsiusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

func drawTable(c celsius) {
	fmt.Printf("| %-9v | %-9.2f |\n", c, celsiusToFahrenheit(c))
}

func main() {
	fmt.Println("========================")
	fmt.Printf("| %-9s | %-9s |\n", "°C", "°F")
	fmt.Println("========================")
	for i := -40; i <= 100; i += 28 {
		drawTable(celsius(i))
	}
}
