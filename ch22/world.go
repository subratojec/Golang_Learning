package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type world struct {
	radius float64
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

var mars = world{radius: 3389.5}

// distance calculation using the spherical lay of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	fmt.Println(spirit)
	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2fkm\n", dist)
}
