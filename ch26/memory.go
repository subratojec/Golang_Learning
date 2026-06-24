package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	fmt.Println(*address)

	fmt.Printf("address is a %T\n", address)

	canada := "Canada"

	var home *string

	fmt.Printf("home is a %T\n", home)

	home = &canada

	// fmt.Printf(home)
	fmt.Println(*home)

	// An asterisk prefixing a type denotes a pointer type, whereas
	// an asterisk prefixing a value denotes a pointer value.
}
