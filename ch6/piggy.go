package main

import (
    "fmt"
    "math/rand"
)

func main() {
    nickels := 0.05
    dimes := 0.10
    quarters := 0.25

    num := rand.Intn(3) + 1

    var amount float64

    for amount <= 20 {
        switch num {
        case 1:
            amount += nickels
        case 2:
            amount += dimes
        case 3:
            amount += quarters
        }

        fmt.Printf("%.2f\n", amount)
    }
}
