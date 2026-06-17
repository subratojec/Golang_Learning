package main

import (
	"fmt"
	"sort"
)

func main() {
	var temperatures = []float64{
		32.5, 33.2, 31.8, 32.0, 33.1, 32.8, 32.4, 32.6, 33.0, 32.9,
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}
	if set[32.0] {
		fmt.Println("set member")
	}
	fmt.Println(set)
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
