package main

import (
	"fmt"
	"math"
)

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}
type gps struct {
	current_location     location
	destination_location location
	world
}
type rover struct {
	gps
}

// Description Method for the location type
func (l location) description(s string) string {
	return fmt.Sprintf("%s: lat=%f, long=%f", s, l.lat, l.long)
}

// Distance Method for the world type
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(math.Pi * p1.lat / 180)
	s2, c2 := math.Sincos(math.Pi * p2.lat / 180)
	clong := math.Cos((math.Pi * (p1.long - p2.long)) / 180)
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// Distance Method for the gps type
func (g gps) distance() float64 {
	return g.world.distance(g.current_location, g.destination_location)
}

// Message Method for the gps type for describing how many kilometers are left to the destination
func (g gps) message() string {
	distance := g.distance()
	return fmt.Sprintf("The Distance left for Destination %v is %.2fkm", g.destination_location, distance)
}

func main() {
	g := gps{
		current_location:     location{lat: -4.5895, long: 137.4417},
		destination_location: location{lat: 4.5, long: 135.9},
		world:                world{radius: 3389.5}, // FOR MARS
	}
	curiosity := rover{g}
	fmt.Println(curiosity.message())
}
