package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := range 5 {
		go sleepyGopher(i, c)
	}
	for i := range 5 {
		gopherID := <-c // Receives a value from a channel
		fmt.Println("Gopher ", gopherID, "has finished sleeping")
		i++
	}
}

func sleepyGopher(id int, c chan int) { //Declare the channel as an argument
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "snore ...")
	c <- id // Sends a value back to main
}
