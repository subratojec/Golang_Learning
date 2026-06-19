package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	for _, row := range u {
		for _, cell := range row {
			if cell {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	for y := range u {
		for x := range u[y] {
			u[y][x] = rand.Float64() < 0.25
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	return x >= 0 &&
		x < width &&
		y >= 0 &&
		y < height &&
		u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {

			// Skip the cell itself
			if dx == 0 && dy == 0 {
				continue
			}

			// Wrap around edges
			nx := (x + dx + width) % width
			ny := (y + dy + height) % height

			if u.Alive(nx, ny) {
				count++
			}
		}
	}

	return count
}

func (u Universe) Next() Universe {
	next := NewUniverse()

	for y := range u {
		for x := range u[y] {

			neighbors := u.Neighbors(x, y)

			if u.Alive(x, y) {
				// Survival rule
				next[y][x] = neighbors == 2 || neighbors == 3
			} else {
				// Birth rule
				next[y][x] = neighbors == 3
			}
		}
	}

	return next
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	u := NewUniverse()
	u.Seed()

	for {
		clearScreen()
		u.Show()

		time.Sleep(100 * time.Millisecond)

		u = u.Next()
	}
}
