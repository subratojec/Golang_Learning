package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("unicorn")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println(len(files))
}
