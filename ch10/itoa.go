package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds."
	fmt.Println(str)

	countd := 9
	strr := fmt.Sprintf("Launch in T minus %v seconds", countd)
	fmt.Println(strr)

	count, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("oh no there is an error %w", err)
	} else {
		fmt.Println(count)
	}
}
