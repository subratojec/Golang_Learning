package main

import "fmt"

/*
 * An anonymous function, also call a function literal in Go,
 * is a function without a name.
 * unlike regular functions, function literals are closures
 * because they keep references to variables in the surrounding scope.
 */
var f = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {
	f()
}
