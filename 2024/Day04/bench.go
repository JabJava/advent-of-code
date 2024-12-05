package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	res, err := Part1()
	if err != nil {
		panic(err)
	}
	endTime := time.Since(startTime)
	fmt.Printf("[PART 1]\nResult: %v \nMilliSeconds: %v\nMicroSeconds: %v\n", res, endTime.Milliseconds(), endTime.Microseconds())

	startTime = time.Now()
	res, err = Part2()
	if err != nil {
		panic(err)
	}
	endTime = time.Since(startTime)
	fmt.Printf("[PART 2]\nResult: %v \nMilliSeconds: %v\nMicroSeconds: %v\n", res, endTime.Milliseconds(), endTime.Microseconds())
}
