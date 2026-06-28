package main

import "fmt"

func sourceGopher(downstream chan string) {
	for _, s := range []string{
		"Hey! My name is Subrato",
		"Hey! My name is Subrato",
		"Good Morning everyone",
		"Hey! My name is Subrato",
	} {
		downstream <- s
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	var prev string
	for item := range upstream {
		if item != prev {
			downstream <- item
			prev = item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for s := range upstream {
		fmt.Println(s)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
