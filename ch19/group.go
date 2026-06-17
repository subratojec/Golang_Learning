package main

import (
	"fmt"
	"math"
)

func main() {
	temp := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64)

	for _, t := range temp {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, temp := range groups {
		fmt.Printf("%v: %v\n", g, temp)
	}

	fmt.Println(math.Trunc(23.0 / 10))
}
