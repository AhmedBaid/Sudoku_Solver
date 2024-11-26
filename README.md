# Sudoku Solver

Ce package implémente un solveur de Sudoku utilisant la méthode de **backtracking**. Il permet de valider, résoudre et afficher des grilles de Sudoku.

## Fonctionnalités

1. **ResolveInput**  
   - Convertit les arguments d'entrée (grille au format texte) en une grille de Sudoku.
   - Retourne une erreur si le format est invalide.

2. **PrintGrid**  
   - Affiche la grille de Sudoku sous une forme lisible.

3. **IsValid**  
   - Vérifie si un chiffre peut être placé dans une position donnée sans violer les règles du Sudoku.

4. **Solve**  
   - Résout la grille de Sudoku en utilisant la méthode de backtracking.

5. **PrintError**  
   - Affiche un message d'erreur en cas de problème.

## Exemple d'utilisation

```go
package main

import (
	"os"
	"sudoku"
)

func main() {
	args := os.Args[1:]
	grid, err := sudoku.ResolveInput(args)
	if err != "" {
		sudoku.PrintError()
		return
	}

	if sudoku.Solve(grid) {
		sudoku.PrintGrid(grid)
	} else {
		sudoku.PrintError()
	}
}
