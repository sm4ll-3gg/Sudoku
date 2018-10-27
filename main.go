package main

import (
	"fmt"

	"github.com/small-egg/sudoku/examples"
)

func main() {
	var field Field
	field.Init(examples.Medium)
	err := field.FindSolution()

	fmt.Printf("\n")
	fmt.Println(field.String())

	if err != nil {
		panic(err)
	}
}
