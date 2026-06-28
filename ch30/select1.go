package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := range 5 {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for range 5 {
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished sleeping.")
		case <-timeout:
			fmt.Println("My patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...", id, "snore ...")
	c <- id
}
