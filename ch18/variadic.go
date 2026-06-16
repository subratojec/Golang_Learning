// To declare a variadic func, us the ellipsis ...
// with the last parameter
package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	// To pass a slice instead of multiple arguments,
	// expand the slice with an ellipsis ...
	planets := []string{"Venus", "Mars", "Jupiter", "Saturn"}
	newPlanets := terraform("new", planets...)
	fmt.Println(newPlanets)
}

/*
 * Three uses for the ellipsis ...
 * 1. Have the Go compiler count the number of elements in a composite literal for an array.
 * 2. Make the last parameter of variadic functions capture zero or more arguments as a slice.
 * 3. Expand the elements of a slice into arguments passed to a function.
 */
