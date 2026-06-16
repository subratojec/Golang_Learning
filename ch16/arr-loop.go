package main

import "fmt"

func main() {
	dwarfs := [5]string{
		"Ceres",
		"Pluto",
		"Hamea",
		"Makemake",
		"Eris",
	}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
