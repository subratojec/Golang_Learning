package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	println(temp)
	temperature["Venus"] = 464
	temperature["Earth"] = 16

	fmt.Println(temperature)

	// moon := temperature["Moon"]
	// fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vC.\n", moon)
	} else {
		fmt.Println("Where is the Moon?")
	}

}
