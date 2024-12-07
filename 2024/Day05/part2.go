package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part2() (int, error) {
	file, err := os.Open("./Day05/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	afterRules := make(map[int][]int)  // elements that comes after keys
	beforeRules := make(map[int][]int) // elements that comes before keys

	scanner := bufio.NewScanner(file)
	pages := make([][]int, 0)
	for scanner.Scan() {
		var line string = scanner.Text()
		if strings.Contains(line, "|") {
			// --------- PARSING RULES ----------
			splitted := strings.Split(line, "|")
			a, _ := strconv.Atoi(splitted[0])
			b, _ := strconv.Atoi(splitted[1])

			//AFTER
			_, ok := afterRules[a]
			if !ok {
				afterRules[a] = make([]int, 0)
			}
			afterRules[a] = append(afterRules[a], b)

			//BEFORE
			_, ok = beforeRules[b]
			if !ok {
				beforeRules[b] = make([]int, 0)
			}
			beforeRules[b] = append(beforeRules[b], a)

		} else if len(line) > 0 {
			// -------------- PARSING PAGES --------------
			pagesString := strings.Split(line, ",")
			pages = append(pages, make([]int, len(pagesString)))
			for i := 0; i < len(pagesString); i++ {
				converted, _ := strconv.Atoi(pagesString[i])
				pages[len(pages)-1][i] = converted
			}
		}
	}
	sum := 0
	for _, page := range pages {
		a, b := checkPages2(page, afterRules, beforeRules)
		if a >= 0 {
			//No good
			for a >= 0 {
				tmp := page[a]
				page[a] = page[b]
				page[b] = tmp
				a, b = checkPages2(page, afterRules, beforeRules)
			}
			middleIndex := len(page) / 2
			sum += page[middleIndex]
		}
	}

	return sum, nil
}

func checkPages2(pages []int, afterRules map[int][]int, beforeRules map[int][]int) (int, int) {
	for i := 0; i < len(pages); i++ {
		setA := afterRules[pages[i]]
		setB := beforeRules[pages[i]]
		if i > 0 && setA != nil {
			a_minus := intersect2(setA, pages[0:i-1])
			if a_minus >= 0 {
				return a_minus, i
			}
		}
		if i < len(pages) && setB != nil {
			a_plus := intersect2(setB, pages[i+1:])
			if a_plus >= 0 {
				return a_plus + i + 1, i
			}
		}
	}
	return -1, -1
}

func intersect2(a []int, b []int) int {
	for i := 0; i < len(a); i++ {
		index := slices.Index(b, a[i])
		if index >= 0 {
			return index
		}
	}
	return -1
}
