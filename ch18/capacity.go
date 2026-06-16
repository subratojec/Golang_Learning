package main

import "fmt"

func main() {
	s := make([]int, 0)

	prevCap := cap(s)

	for i := 0; i < 2000; i++ {
		s = append(s, i)

		if cap(s) != prevCap {
			fmt.Printf("len=%4d cap=%4d\n", len(s), cap(s))
			prevCap = cap(s)
		}
	}
}
