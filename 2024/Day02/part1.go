package main

import (
	"bufio"
	"math"
	"os"
	"strings"

	// "strings"

	"strconv"
)

func Part1() (int, error) {
	file, err := os.Open("./Day02/input.txt")
	if err != nil {
		return -1, err
	}
	defer file.Close()
	goodL := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var couple string = scanner.Text()
		levels := strings.Split(couple, " ")
		lstLevel, err := strconv.Atoi(levels[0])
		if err != nil {
			return -1, err
		}
		isgood := true
		secondL, err := strconv.Atoi(levels[1])
		if err != nil {
			return -1, err
		}
		isIncrease := (float64(lstLevel) - float64(secondL)) < 0
		//fmt.Printf("Level: %v | Incr: %v\n", lstLevel, isIncrease)
		for i := 1; i < len(levels); i++ {
			lvl, err := strconv.Atoi(levels[i])
			if err != nil {
				return -1, err
			}
			signedDelta := float64(lstLevel) - float64(lvl)
			//Rule1
			if signedDelta == 0 || (signedDelta < 0 && !isIncrease) || (signedDelta > 0 && isIncrease) {
				isgood = false
				break
			}
			//Rule2
			delta := math.Abs(signedDelta)
			//fmt.Printf("\t Rule2 => Delta: %v\n", lstLevel, isIncrease)
			if delta == 0 || delta > 3 {
				isgood = false
				break
			}
			lstLevel = lvl
		}
		if isgood {
			goodL += 1
		}
	}
	return goodL, nil
}
