package sudoku

import "github.com/01-edu/z01"

// ResolveInput prend en entrée les arguments et les convertit en une grille de sudoku.
func ResolveInput(args []string) ([][]int, string) {
	grid := make([][]int, 9)
	for i, row := range args {
		if len(row) != 9 {
			return nil, "Error"
		}
		grid[i] = make([]int, 9)
		for j, r := range row {
			if r == '.' {
				grid[i][j] = 0
			} else if r >= '1' && r <= '9' {
				grid[i][j] = int(r - '0')
			} else {
				return nil, "Error"
			}
		}
	}
	return grid, ""
}

// PrintGrid affiche la grille de sudoku.
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for j, val := range row {
			z01.PrintRune(rune(val + '0'))
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

// IsValid vérifie si un nombre peut être placé à une position donnée dans la grille.
func IsValid(grid [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

// Solve résout le sudoku en utilisant la méthode de backtracking.
func Solve(grid [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if IsValid(grid, row, col, num) {
						grid[row][col] = num
						if Solve(grid) {
							return true
						}
						grid[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// PrintError affiche un message d'erreur.
func PrintError() {
	for _, r := range "Error\n" {
		z01.PrintRune(r)
	}
}
