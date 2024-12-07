package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	// "strings"
)

func Part1() (int, error) {
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
		if checkPages(page, afterRules, beforeRules) {
			middleIndex := len(page) / 2
			//fmt.Printf("Adding %v\n", page[middleIndex])
			sum += page[middleIndex]
		}
	}

	return sum, nil
}

func checkPages(pages []int, afterRules map[int][]int, beforeRules map[int][]int) bool {
	for i := 0; i < len(pages); i++ {
		setA, _ := afterRules[pages[i]]
		setB, _ := beforeRules[pages[i]]
		if i > 0 && setA != nil {
			a_minus := intersect(setA, pages[i-1:i])
			if a_minus {
				return false
			}
		}
		if i < len(pages) && setB != nil {
			a_plus := intersect(setB, pages[i+1:])
			if a_plus {
				return false
			}
		}
	}
	return true
}

func intersect(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if slices.Contains(b, a[i]) {
			return true
		}
	}
	return false
}
