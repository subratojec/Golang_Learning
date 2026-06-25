package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name     string
	lefthand *item
}

func (c *character) pickup(i *item) {
	if c.lefthand == nil {
		fmt.Printf("%s holds nothing\n", c.name)
	}
	c.lefthand = i
	fmt.Printf("%s holds %s\n", c.name, i.name)
}

func (c *character) give(to *character) {
	if c.lefthand == nil {
		fmt.Printf("%s holds nothing\n", c.name)
		return
	}
	item := c.lefthand
	c.lefthand = nil
	to.pickup(item)
	fmt.Printf("%s gave %s to %s.\n", c.name, item.name, to.name)
}

func main() {
	sword := item{name: "sword"}
	subrato := character{name: "subrato"}
	sani := character{name: "sani"}
	subrato.pickup(&sword)
	subrato.give(&sani)
}
