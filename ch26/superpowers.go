package main

import "fmt"

func main() {
	superpowers := &[3]string{"Flight", "Invisibility", "Super Strenght"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])
	fmt.Println(superpowers[1:])
}

// Arrays also provide automaitc dereferencing
// Composite literals for slices and maps can also be prefixed with the
// address operator (&), but there's no automatic dereferencing
