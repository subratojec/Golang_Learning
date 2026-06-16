package main

import "fmt"

type Planets []string

func (p Planets) terraform() {
	for i := range len(p) {
		p[i] = "New " + p[i]
	}
}

func main() {
	planets := Planets{"Mars", "Uranus", "Neptune"}
	// for i := range len(planets) {
	// 	planets[i] = fmt.Sprintf("`New %s`", planets[i])
	// }
	// fmt.Println(planets)
	planets.terraform()
	fmt.Println("`New` is added in the prepending.")
	fmt.Println(planets)

}
