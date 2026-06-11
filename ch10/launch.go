package main
import "fmt"

func main() {
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	var yesNO string
	if launch {
		yesNO = "yes"
	} else {
		yesNO = "no"
	}
	fmt.Println("Ready for launch:", yesNO)
}
