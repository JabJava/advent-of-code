package main

import (
	"bufio"
	"math"
	"os"
	"strings"
	// "strings"
)

func Part1() (int, error) {
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
			if grid[y][x] != "X" {
				continue
			}
			if checkXMAS(grid, -1, 0, x, y) {
				sum++
			}
			if checkXMAS(grid, 1, 0, x, y) {
				sum++
			}
			if checkXMAS(grid, 0, -1, x, y) {
				sum++
			}
			if checkXMAS(grid, 0, 1, x, y) {
				sum++
			}

			if checkXMAS(grid, 1, 1, x, y) {
				sum++
			}

			if checkXMAS(grid, -1, -1, x, y) {
				sum++
			}

			if checkXMAS(grid, -1, 1, x, y) {
				sum++
			}

			if checkXMAS(grid, 1, -1, x, y) {
				sum++
			}

		}
	}
	return sum, nil
}
func checkXMAS(grid [][]string, x_off int, y_off int, x int, y int) bool {
	acc := ""
	xoffset := x_off * 3
	yoffset := y_off * 3
	if yoffset != 0 {
		for yy := Min(yoffset, 0); yy < Min(yoffset, 0)+4; yy++ {
			if xoffset == 0 {
				if y+yy < 0 {
					return false
				}
				if y+yy >= len(grid) {
					return false
				}
				acc += grid[y+yy][x]
			} else {
				for xx := Min(xoffset, 0); xx < Min(xoffset, 0)+4; xx++ {
					if math.Abs(float64(xx)) == math.Abs(float64(yy)) {
						if x+xx < 0 {
							return false
						}
						if x+xx >= len(grid[0]) {
							return false
						}
						if y+yy < 0 {
							return false
						}
						if y+yy >= len(grid) {
							return false
						}
						acc += grid[y+yy][x+xx]
					}
				}
			}
		}
		return acc == "XMAS" || acc == "SAMX"
	} else {
		for xx := Min(xoffset, 0); xx < Min(xoffset, 0)+4; xx++ {
			if x+xx < 0 {
				return false
			}
			if x+xx >= len(grid[0]) {
				return false
			}
			acc += grid[y][x+xx]
		}
		return acc == "XMAS" || acc == "SAMX"
	}
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
