package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2() (int, error) {
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
		var isGood bool = CheckLevels(levels)
		if isGood {
			goodL++
		} else {
			//Bruteforce because i'm dumb :(
			for i := 0; i < len(levels); i++ {
				cpy := make([]string, len(levels))
				copy(cpy, levels)
				var sliced []string
				if i == 0 {
					sliced = cpy[1:]
				} else if i == len(cpy)-1 {
					sliced = cpy[0 : len(cpy)-1]
				} else {
					sliced = append(cpy[0:i], cpy[i+1:]...)
				}
				if CheckLevels(sliced) {
					goodL++
					isGood = true
					break
				}
			}
		}
	}
	return goodL, nil
}
func CheckLevels(levels []string) bool {
	lstLevel, err := strconv.Atoi(levels[0])
	if err != nil {
		panic(err)
	}
	isgood := true
	secondL, err := strconv.Atoi(levels[1])
	if err != nil {
		panic(err)
	}
	isIncrease := (float64(lstLevel) - float64(secondL)) < 0
	//fmt.Printf("Level: %v | Incr: %v\n", lstLevel, isIncrease)
	for i := 1; i < len(levels); i++ {
		lvl, err := strconv.Atoi(levels[i])
		if err != nil {
			panic(err)
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
	return isgood
}
