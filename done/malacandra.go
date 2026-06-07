package main

import (
	"fmt"
)

func main() {
	const hour = 28 * 24
	var distance = 56000000 //km
	var time = distance / hour
	fmt.Println(time,"km/h")
}
