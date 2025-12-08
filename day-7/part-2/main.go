package main

import (
	"fmt"
	"os"
	"strings"
)

const numberPlaceholder = '|'

func main() {
	lines := readInputLines("day-7/input.txt")
	count := countTimelines(lines)
	fmt.Println(count)
}

func countTimelines(lines []string) int {
	timelineCount := 0
	numbers := make(map[string]int)
	for row, line := range lines[:len(lines)-1] {
		if row == 0 {
			continue
		}
		for col, char := range line {
			topCenterChar := lines[row-1][col]
			num := 0

			if char == '^' {
				continue
			}

			if topCenterChar == 'S' {
				num = 1
			}

			if topCenterChar == numberPlaceholder {
				num += getNumber(numbers, row-1, col)
			}

			if col > 0 && lines[row][col-1] == '^' && lines[row-1][col-1] == numberPlaceholder {
				num += getNumber(numbers, row-1, col-1)
			}

			if col < len(line)-1 && lines[row][col+1] == '^' && lines[row-1][col+1] == numberPlaceholder {
				num += getNumber(numbers, row-1, col+1)
			}

			if num > 0 {
				lines[row] = lines[row][:col] + string(numberPlaceholder) + lines[row][col+1:]
				setNumber(numbers, row, col, num)
			}
		}

		if row == len(lines)-2 {
			for col, char := range lines[row] {
				if char == numberPlaceholder {
					timelineCount += getNumber(numbers, row, col)
				}
			}
		}
	}
	return timelineCount
}

func getNumber(numbers map[string]int, row, col int) int {
	return numbers[fmt.Sprintf("%d,%d", row, col)]
}

func setNumber(numbers map[string]int, row, col, value int) {
	numbers[fmt.Sprintf("%d,%d", row, col)] = value
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
