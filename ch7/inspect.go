package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type: %T for %v\n", year,year)
	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)
	text := "Subrato Singh"
	num := 42
	word := true
	fmt.Printf("Type %T for %v\n", text, text)
	fmt.Printf("Type %T for %v\n", num, num)
	fmt.Printf("Type %T for %v\n", word, word)

	// var red, green, blue unit8 = 0,141,213
	// var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	// fmt.Printf("%x %x %x\n", red, green, blue)

	// fmt.Printf("color: #%02x%02x%02x\n", red, green, blue)
	var red uint8 = 255
	red++
	fmt.Println(red)
	var number int8 = 127
	number++
	fmt.Println(number)

	// When the range is exceeded, interger types in Go wrap around.
}
