// METHODS EXPRESS THE BEHAVIOR A TYPE PROVIDES, SO
// INTERFACES ARE DECLARED WITH A SET OF METHODS THAT A TYPE MUST STATISFY.
package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

// The variable t can hold any value of any type that statifies the interface.
// More specfically, a type will specify the interface if it declares a method named talk that accepts no
// arguments and returns a string.

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

func main() {
	// var t interface {
	// 	talk() string
	// }
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(3))

	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}
