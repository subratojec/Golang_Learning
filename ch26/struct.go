package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

// It isn't necessary to dereference structures to access their fields.
func main() {
	subrato := &person{
		name: "Subrato",
		age:  22,
	}
	subrato.superpower = "flying"
	fmt.Printf("%+v\n", subrato)
}
