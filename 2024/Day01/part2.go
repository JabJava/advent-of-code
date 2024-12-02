package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func Part2() (int, error) {
	r := regexp.MustCompile(`(\d+)\s+(\d+)`)
	file, err := os.Open("./Day01/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	left := make(map[int]int)
	right := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var couple string = scanner.Text()
		matches := r.FindStringSubmatch(couple)
		leftN, _ := strconv.Atoi(matches[1])
		rightN, _ := strconv.Atoi(matches[2])

		cnt, ok := left[leftN]
		if !ok {
			left[leftN] = 1
		} else {
			left[leftN] = cnt + 1
		}

		cnt, ok = right[rightN]
		if !ok {
			right[rightN] = 1
		} else {
			right[rightN] = cnt + 1
		}
	}
	sum := 0
	for key, val := range left {
		sum += (key * right[key]) * val
	}

	return sum, nil
}
