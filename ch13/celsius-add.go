package main

import "fmt"

func main() {
	type celsius float64
	const degrees = 20
	var temp celsius = degrees
	temp += 10
	fmt.Println(temp)

	var warmUp float64 = 10
	temp += celsius(warmUp)
	fmt.Println(temp)

	type fahrenheit float64

	var c celsius = 20
	var f fahrenheit = 20

	if c == f {
		fmt.Println("celsius and fahrenheit are equal")
	}
	c += f
	fmt.Println(c)
}
