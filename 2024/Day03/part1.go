package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// "strings"

func Part1() (int, error) {

	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	file, err := os.Open("./Day03/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var instruction string = scanner.Text()
		matches := r.FindAllStringSubmatch(instruction, -1)
		for _, v := range matches {
			sum += mul(v)
		}
	}
	return sum, nil
}

func mul(lst []string) int {
	a, err := strconv.Atoi(lst[1])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(lst[2])
	if err != nil {
		panic(err)
	}

	return a * b
}
