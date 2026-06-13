package main

import "fmt"

func main() {
	// type keyword declares a new type with
	//  a name and underlying type
	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature)

}
