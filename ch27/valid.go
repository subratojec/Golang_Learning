package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) string() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {
	n := newNumber(42)
	fmt.Println(n.string())

	e := number{}
	fmt.Println(e.string())
}
