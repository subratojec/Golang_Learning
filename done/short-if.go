package main
import (
	"fmt"
	"math/rand"
)
func main() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("space adventures")
	} else if num == 1 {
		fmt.Println("Space X")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
