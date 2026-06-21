package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low float64
}

// type celsius float64

type location struct {
	lat, long float64
}

// average.go: average method using the temperature type.
func (t temperature) average() float64 {
	return (t.high + t.low) / 2
}

func (r report) average() float64 {
	return r.temperature.average()
}

func main() {
	// bradbury := location{lat: 40.8244, long: -111.7638}
	// t := temperature{high: -1.0, low: -78.0}
	// report := report{sol: 15, temperature: t, location: bradbury}
	// fmt.Printf("%+v\n", report)
	// fmt.Printf("a balmy %v C\n", report.temperature.high)
	t := temperature{high: -1.0, low: -78.0}
	// fmt.Printf("average %vC\n", t.average())

	report := report{sol: 15, temperature: t}
	fmt.Printf("average %vC\n", report.temperature.average())
	fmt.Printf("average %vC\n", report.average())

}
