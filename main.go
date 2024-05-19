//nolint:errcheck
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func checkParam(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, arg := range args {
		if len(arg) != 9 {
			return false
		}
		for _, char := range arg {
			if char != '.' && (char < '1' || char > '9') {
				return false
			}
		}
	}
	return true
}

func printGrid(grid [][]byte) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if x > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(grid[y][x]))
		}
		z01.PrintRune('\n')
	}
}

func printError() {
	errorMsg := "Error\n"
	for _, char := range errorMsg {
		z01.PrintRune(char)
	}
}

func checkSudoku(deep int, grid [][]byte) bool {
	x := deep % 9
	y := deep / 9
	// Check row
	for i := 0; i < 9; i++ {
		if i != x && grid[y][i] == grid[y][x] {
			return false
		}
	}
	// Check column
	for j := 0; j < 9; j++ {
		if j != y && grid[j][x] == grid[y][x] {
			return false
		}
	}
	// Check 3x3 sub-box
	caseX := x / 3
	caseY := y / 3
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			xx := caseX*3 + i
			yy := caseY*3 + j
			if (xx != x || yy != y) && grid[yy][xx] == grid[y][x] {
				return false
			}
		}
	}
	return true
}

func backTrack(deep int, grid [][]byte) bool {
	if deep == 81 {
		printGrid(grid)
		return true
	}
	x := deep % 9
	y := deep / 9
	if grid[y][x] != '.' {
		if checkSudoku(deep, grid) && backTrack(deep+1, grid) {
			return true
		}
	} else {
		for i := '1'; i <= '9'; i++ {
			grid[y][x] = byte(i)
			if checkSudoku(deep, grid) && backTrack(deep+1, grid) {
				return true
			}
		}
		grid[y][x] = '.'
	}
	return false
}

func resolveSudoku(args []string) bool {
	var grid [][]byte
	for _, arg := range args {
		grid = append(grid, []byte(arg))
	}
	return backTrack(0, grid)
}

func main() {
	args := os.Args
	if !checkParam(args[1:]) || !resolveSudoku(args[1:]) {
		printError()
	}
}

// test
