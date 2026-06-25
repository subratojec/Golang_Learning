package main

import (
	"fmt"
	"sort"
)

var fn func(a, b int) int

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func main() {
	food := []string{"onion", "carrot", "celery", "appless"}
	// sortStrings(food, nil)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) })
	fmt.Println(food)
}
