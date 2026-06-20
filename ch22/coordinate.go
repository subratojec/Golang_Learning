// This how we create constructors in Golang
// newLocation is a constructor that returns a location struct
// new is a name convention for constructors in Golang

package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// func main() {
// 	//Bradbury Landing
// 	lat := coordinate{4, 35, 22.2, 'S'}
// 	long := coordinate{70, 46, 5.9, 'E'}

// 	fmt.Println(lat.decimal(), long.decimal())

func main() {
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{70, 46, 5.9, 'E'})
	fmt.Println(curiosity)
}
