package main

import (
	"fmt"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

const joltageLength = 12

func main() {
	totalJoltage := 0

	util.ScanInput("day-3/input.txt", func(input string) {
		bank, _ := stringToDigits(input)
		totalJoltage += findLargestJoltage(bank)
	})

	fmt.Println(totalJoltage)
}

func findLargestJoltage(bank []int) int {
	joltageDigits := make([]int, joltageLength)
	startIndex := 0

	for index := range joltageLength {
		for j := startIndex; j <= len(bank)-joltageLength+index; j++ {
			if joltageDigits[index] < bank[j] {
				joltageDigits[index] = bank[j]
				startIndex = j + 1
			}
		}
	}

	return digitsToNumber(joltageDigits)
}

func stringToDigits(str string) ([]int, error) {
	digits := make([]int, len(str))
	for i, ch := range str {
		if ch < '0' || ch > '9' {
			return nil, fmt.Errorf("non-digit at index %d", i)
		}
		digits[i] = int(ch - '0')
	}
	return digits, nil
}

func digitsToNumber(digits []int) int {
	number := 0
	for _, digit := range digits {
		number = number*10 + digit
	}
	return number
}
