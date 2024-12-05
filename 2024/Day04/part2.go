package main

import (
	"bufio"
	"os"
	"strings"
)

func Part2() (int, error) {
	file, err := os.Open("./Day04/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)
	for scanner.Scan() {
		var row string = scanner.Text()
		grid = append(grid, strings.Split(row, ""))
	}
	sum := 0
	for y := 0; y < len(grid); y++ {
		row := grid[y]
		for x := 0; x < len(row); x++ {
			if grid[y][x] != "A" {
				continue
			}
			diag1 := checkXMAS2(grid, true, x, y)
			diag2 := checkXMAS2(grid, false, x, y)
			if diag1 && diag2 {
				sum++
			}
		}
	}
	return sum, nil
}

func checkXMAS2(grid [][]string, dir bool, x int, y int) bool {
	acc := ""
	if dir {
		for offset := -1; offset < 2; offset++ {
			yy := y + offset
			xx := x + offset
			if xx < 0 {
				return false
			}
			if xx >= len(grid[0]) {
				return false
			}
			if yy < 0 {
				return false
			}
			if yy >= len(grid) {
				return false
			}
			acc += grid[yy][xx]
		}
	} else {
		for offset := -1; offset < 2; offset++ {
			yy := y + (offset * -1)
			xx := x + offset
			if xx < 0 {
				return false
			}
			if xx >= len(grid[0]) {
				return false
			}
			if yy < 0 {
				return false
			}
			if yy >= len(grid) {
				return false
			}
			acc += grid[yy][xx]
		}
	}
	return acc == "SAM" || acc == "MAS"

}
