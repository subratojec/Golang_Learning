package main

import "fmt"

func main() {
	temp := []float64{
		-28.0, 32.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	freq := make(map[float64]int)

	for _, t := range temp {
		freq[t]++
	}

	for t, num := range freq {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}
