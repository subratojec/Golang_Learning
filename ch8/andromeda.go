package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("41300000000000", 10)
	fmt.Println("Andromeda is", distance, "km away.")

	seconds := new(big.Int)
	seconds.Div(distance,lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("That is", days, "days of travel at light speed")
}
