package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector zz9",
		"Mars":  "Sector zz9",
	}

	planetsMark2 := planets

	planets["Earth"] = "Whoops"

	fmt.Println(planets)
	fmt.Println(planetsMark2)

	delete(planets, "Earth")

	fmt.Println(planetsMark2)
}
