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

func newLocation(lat, long coordinate) location {
	loc := location{
		lat:  lat.decimal(),
		long: long.decimal(),
	}
	return loc
}

// distance calculation using the spherical lay of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(math.Pi * p1.lat / 180)
	s2, c2 := math.Sincos(math.Pi * p2.lat / 180)
	clong := math.Cos((math.Pi * (p1.long - p2.long)) / 180)
	return w.radius * math.Acos(s1*s2+c1*c2*clong) // Corrected this line
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

type Rover struct {
	Name        string
	LandingSite string
	Latitude    coordinate
	Longitude   coordinate
	location    location
}

var mars = world{radius: 3389.5}
var earth = world{radius: 6371.0}

func main() {
	rovers := []Rover{
		{
			Name:        "Spirit",
			LandingSite: "Columbia Memorial Station",
			Latitude:    coordinate{14, 34, 6.2, 'S'},
			Longitude:   coordinate{175, 221.5, 0, 'E'},
		},
		{
			Name:        "Opportunity",
			LandingSite: "Challenger Memorial Station",
			Latitude:    coordinate{-1, 56, 46.3, 'S'},
			Longitude:   coordinate{354, 28, 24.2, 'E'},
		},
		{
			Name:        "Curiosity",
			LandingSite: "Bradbury Landing",
			Latitude:    coordinate{4, 30, 0.0, 'N'},
			Longitude:   coordinate{137, 26, 30.1, 'E'},
		},
		{
			Name:        "InSight",
			LandingSite: "Elysium Planitia",
			Latitude:    coordinate{4, 30, 0.0, 'N'},
			Longitude:   coordinate{135, 54, 0, 'E'},
		},
	}
	for i := range rovers {
		rovers[i].location = newLocation(
			rovers[i].Latitude,
			rovers[i].Longitude,
		)
		fmt.Println(rovers[i].location)
	}

	Farthest := 0.0
	nearest := 0.0
	for i := 0; i < len(rovers); i++ {
		for j := i + 1; j < len(rovers); j++ {
			d := mars.distance(rovers[i].location, rovers[j].location)
			fmt.Printf("%v vs %v: %.2f\n", rovers[i].Name, rovers[j].Name, d)

			if d > Farthest {
				Farthest = d
			}
			if d < nearest || nearest == 0 {
				nearest = d
			}
		}
	}
	fmt.Printf("Max distance: %0.2f\n", Farthest)
	fmt.Printf("Min distance: %0.2f\n", nearest)

	london := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 0.8, 0, 'W'})
	france := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})

	fmt.Println("Distance between London and France: ", earth.distance(london, france), "km")

}
