package main

import "fmt"

type pos struct {
	x, y int
}

func (p *pos) moveUp() {
	p.y++
}

func (p *pos) moveDown() {
	p.y--
}

func (p *pos) moveLeft() {
	p.x--
}

func (p *pos) moveRight() {
	p.x++
}

func (p *pos) location() (int, int) {
	return p.x, p.y
}

type turtle struct {
	pos
}

func main() {
	t := turtle{
		pos: pos{
			x: -5,
			y: 0,
		},
	}
	t.moveUp()
	t.moveRight()
	t.moveDown()
	t.moveLeft()
	t.moveUp()
	x, y := t.location()
	fmt.Println(x, y)
}
