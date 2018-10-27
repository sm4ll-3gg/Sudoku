package main

import (
	"fmt"

	"github.com/small-egg/sudoku/examples"
)

func main() {
	var field Field
	field.Init(examples.Simpliest)
	err := field.FindSolution()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")
	fmt.Println(field.String())
}
