package main

import (
	"fmt"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

func main() {
	totalJoltage := 0

	util.ScanInput("day-3/input.txt", func(input string) {
		bank, _ := digitsToSlice(input)
		totalJoltage += findLargestJoltage(bank)
	})

	fmt.Println(totalJoltage)
}

func findLargestJoltage(bank []int) int {
	firstJoltageDigit, secondJoltageDigit := 0, 0

	for index, joltage := range bank {
		isLastDigit := index == len(bank)-1
		if firstJoltageDigit < joltage && !isLastDigit {
			firstJoltageDigit = joltage
			secondJoltageDigit = 0
		} else if secondJoltageDigit < joltage {
			secondJoltageDigit = joltage
		}
	}

	return firstJoltageDigit*10 + secondJoltageDigit
}

func digitsToSlice(s string) ([]int, error) {
	out := make([]int, len(s))
	for i, ch := range s {
		if ch < '0' || ch > '9' {
			return nil, fmt.Errorf("non-digit at index %d", i)
		}
		out[i] = int(ch - '0')
	}
	return out, nil
}
