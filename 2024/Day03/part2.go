package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part2() (int, error) {
	r := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)
	file, err := os.Open("./Day03/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	txt := ""
	for scanner.Scan() {
		var instruction string = scanner.Text()
		txt += instruction

	}
	matches := r.FindAllStringSubmatch(txt, -1)
	enabled := true
	for _, v := range matches {
		if strings.HasPrefix(v[0], "mul(") {
			if enabled {
				sum += mul2(v)
			}
		} else if strings.HasPrefix(v[0], "don't(") {
			enabled = false
		} else if strings.HasPrefix(v[0], "do(") {
			enabled = true
		}
	}
	return sum, nil
}

func mul2(lst []string) int {
	a, err := strconv.Atoi(lst[2])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(lst[3])
	if err != nil {
		panic(err)
	}
	return a * b
}
