package adventofcode

import (
	"fmt"
	"os"
	"strings"
)

func isValidRC(r int, c int, rows int, cols int) bool {
	return 0 <= r && r < rows && 0 <= c && c < cols
}

func getCrossMASCount(grid []string) int {
	rows := len(grid)
	cols := len(grid[0]) - 1
	count := 0
	dirs := [][2]int{{-1, -1}, {-1, 1},  {1, 1}, {1, -1}}
	for r := range rows {
		for c := range cols {
			if grid[r][c] != 'A' {
				continue
			}

			// Check all 4 diagonals and see if letters are present
			M_count := 0
			S_count := 0
			diag := make([]string, 0)
			for _, val := range dirs {
				new_r := r + val[0]
				new_c := c + val[1]
				if isValidRC(new_r, new_c, rows, cols) {
					if grid[new_r][new_c] == 'M' {
						M_count += 1
						diag = append(diag, "M")
					} else if grid[new_r][new_c] == 'S' {
						S_count += 1
						diag = append(diag, "S")
					}
				}
			}
			if len(diag) == 4 && M_count == 2 && S_count == 2 {
				// Check if values are consecutive.
				for idx := range 3 {
					if diag[idx] == diag[idx + 1] {
						count += 1
						break
					}
				}
			}
			
		}
	}
	return count
}

func getXMASCount(grid []string) int {
	rows := len(grid)
	cols := len(grid[0]) - 1
	count := 0
	dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for r := range rows {
		for c := range cols {
			if grid[r][c] != 'X' {
				continue
			}
			// Check all 9 directions and 4 letters
			for _, val := range dirs {
				new_r1 := r + val[0]
				new_c1 := c + val[1]
				new_r2 := new_r1 + val[0]
				new_c2 := new_c1 + val[1]
				new_r3 := new_r2 + val[0]
				new_c3 := new_c2 + val[1]
				if isValidRC(new_r1, new_c1, rows, cols) && isValidRC(new_r2, new_c2, rows, cols) && isValidRC(new_r3, new_c3, rows, cols) {
					if grid[new_r1][new_c1] == 'M' && grid[new_r2][new_c2] == 'A' && grid[new_r3][new_c3] == 'S' {
						count += 1
					}
				}
			}
		}
	}
	return count
}

func adventDay4() {
	dat, err := os.ReadFile("inputs/input_day4.txt")
	if err != nil {
		fmt.Println("Error Reading Input :", err)
		return
	}
	input := strings.Split(string(dat), "\n")
	fmt.Println("Cross MAS Count :", getCrossMASCount(input))
}