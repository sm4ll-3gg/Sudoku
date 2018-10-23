package main

import (
	"fmt"

	"github.com/small-egg/sudoku/examples"
)

func main() {
	var field Field
	field.Init(examples.Medium)

	fmt.Printf("\n")
	fmt.Println(field.String())
}
