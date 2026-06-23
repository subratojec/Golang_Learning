package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, lon coordinate
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%v\"%c", c.d, c.m, c.s, c.h)
}

func (l location) String() string {
	return fmt.Sprintf("%v,%v", l.lat, l.lon)
}

func main() {
	elysium := location{
		lat: coordinate{4, 30, 0.0, 'N'},
		lon: coordinate{7, 30, 0.0, 'W'},
	}

	fmt.Println("Elysium Planitia is at", elysium)
}
