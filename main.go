package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	m := NewMachine(string(code), os.Stdin, os.Stdout)
	m.Execute()
}
