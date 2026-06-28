package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go sleepyGopher()
// 	time.Sleep(4 * time.Second)
// }

func main() {
	for i := range 5 {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}
func sleepyGopher(i int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...", i)
}
