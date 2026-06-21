package main

import "fmt"

// TYPES
type temperature struct {
	high, low float64
}

type location struct {
	lat, long float64
}

type sol int

type report struct {
	// sol int
	sol
	temperature
	location
}

// METHODS

func (t temperature) average() float64 {
	return (t.high + t.low) / 2
}

func (l location) days(l2 location) int {
	// To-do: complicated distance calculation
	return 5
}

// Any methods declared on the sol type can be accessed through the sol field or through the report type:
func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

// MAIN FUNCTION
func main() {
	// report := report{
	// 	sol:         15,
	// 	location:    location{-4.5895, 137.4417},
	// 	temperature: temperature{high: -1.0, low: -78.0},
	// }
	// fmt.Printf("average: %v\n", report.average())
	// fmt.Printf("%v \n", report.high)
	// report.high = 32
	// fmt.Printf("%v\n", report.temperature.high)

	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))

	d := report.days(1446)
	fmt.Println(d) // ambiguous selector report.days, Go compiler doesn't know if it should forward the call to the method on sol or the method on location, so its gives an error.
}
