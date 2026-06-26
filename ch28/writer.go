package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check erros, handle them gracefully.")
	sw.writeln("Don't Panic")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The Bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Documentation is for users.")

	return sw.err
}

func main() {
	err := proverbs("go_proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File Created")
}
