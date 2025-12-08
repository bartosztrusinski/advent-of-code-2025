package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInputLines("day-7/input.txt")
	count := countSplits(lines)
	fmt.Println(count)
}

func countSplits(lines []string) int {
	count := 0
	for i, line := range lines[:len(lines)-1] {
		for j, char := range line {
			if char == 'S' || char == '|' {
				nextLine := lines[i+1]
				if nextLine[j] == '^' {
					lines[i+1] = nextLine[:j-1] + "|^|" + nextLine[j+2:]
					count++
				} else {
					lines[i+1] = nextLine[:j] + "|" + nextLine[j+1:]
				}
			}
		}
	}
	return count
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
