package main

import (
	"slices"
)

func main() {
	setA := []int{1, 2, 3}
	pages := []int{1, 2, 3, 4, 5, 6, 7, 8}

}

func check(pages []int, setA []int, setB []int) bool {
	for i := 0; i < len(pages); i++ {
		if i > 0 {
			a_minus := intersect(setA, pages[i-1:i])
			if a_minus {
				return false
			}
		}
		if i < len(pages) {
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
