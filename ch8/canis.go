package main

import (
	"fmt"
)

func main() {
	const lightSpeed = 299792
	const distance = 236000000000000000
	const secondsPerYear = 31536000

	const year = distance / lightSpeed / secondsPerYear
	fmt.Println("Canis Major", year ,"light years away")
}
