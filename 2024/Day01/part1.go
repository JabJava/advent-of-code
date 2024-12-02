package main

import (
	"bufio"
	"math"
	"os"

	// "strings"

	"regexp"
	"strconv"
)

func Part1() (int, error) {
	r := regexp.MustCompile(`(\d+)\s+(\d+)`)
	file, err := os.Open("./Day01/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var couple string = scanner.Text()
		matches := r.FindStringSubmatch(couple)
		leftN, _ := strconv.Atoi(matches[1])
		rightN, _ := strconv.Atoi(matches[2])
		left = append(left, leftN)
		right = append(right, rightN)
		bubbleUp(left)
		bubbleUp(right)
	}
	sum := 0

	for i := 0; i < len(left); i++ {
		delta := math.Abs(float64(left[i]) - float64(right[i]))
		sum += int(delta)
	}

	return sum, nil
}

func bubbleUp(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			tmp := arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
		} else {
			return
		}
	}
}
