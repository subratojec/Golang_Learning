package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	name         string
	landing_site string
	lat          float64
	long         float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1.0
	}
	return sign * (c.d + c.m/60 + c.s/3600)

}

func newLocation(name string, landing_site string, lat, long coordinate) location {
	return location{
		name:         name,
		landing_site: landing_site,
		lat:          lat.decimal(),
		long:         long.decimal(),
	}
}

func main() {
	spirit := newLocation("Spirit", "Columbia Memorial Station", coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	fmt.Println(spirit)
}
