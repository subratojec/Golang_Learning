package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Makemake", "Haumea", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Salacia", "Quaoar")
	fmt.Println(dwarfs)
	fmt.Println(len(dwarfs))
}
