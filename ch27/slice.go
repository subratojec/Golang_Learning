package main

import "fmt"

func mirepoix(indegredients []string) []string {
	return append(indegredients, "onions", "carrots", "celery")
}

func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)

}
