package main
import "fmt"
func main() {
	var weight = 149.0
	weight = weight * 0.3784
	weight *= 0.3783
	fmt.Println(weight)

	var age = 41
	age = age + 1
	age += 1
	age ++
	fmt.Println(age)
}
