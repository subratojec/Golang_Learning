package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// closure : A function that remembers things from where it was created.
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

// func main() {
// 	offset := kelvin(5)
// 	sensor := calibrate(realSensor, offset)
// 	fmt.Println(sensor())
// 	offset = 100
// 	fmt.Println(sensor())
// 	fmt.Println(sensor())
// }

func main() {
	sensor := calibrate(fakeSensor, 5)

	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
