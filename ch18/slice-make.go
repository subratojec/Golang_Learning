/*
 * The make function specifies both the length and cap of the slice.
 * elements can be appended before slice runs out of capacity, causing
 * append to allocate a new array.
 */

package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs, cap(dwarfs))
}
