package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInputLines("day-7/input.txt")
	count := 0
	for i, line := range lines {
		if i == len(lines)-1 {
			break
		}
		for j, char := range line {
			if char == 'S' || char == '|' {
				if lines[i+1][j] == '^' {
					lines[i+1] = lines[i+1][:j-1] + "|^|" + lines[i+1][j+2:]
					count++
				} else {
					lines[i+1] = lines[i+1][:j] + "|" + lines[i+1][j+1:]
				}
			}
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println(count)
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
