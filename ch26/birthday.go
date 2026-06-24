package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	subrato := person{
		name:       "Subrato",
		superpower: "Imagination",
		age:        22,
	}
	birthday(&subrato)
	fmt.Printf("%+v\n", subrato)
}
