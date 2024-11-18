package main

import (
	"os"

	"sudoku" // Importer le package sudoku
)

func main() {
	if len(os.Args[1:]) != 9 {
		sudoku.PrintError()
		return
	}

	grid, err := sudoku.ResolveInput(os.Args[1:])
	if err == "Error" {
		sudoku.PrintError()
		return
	}

	if sudoku.Solve(grid) {
		sudoku.PrintGrid(grid)
	} else {
		sudoku.PrintError()
	}
}
