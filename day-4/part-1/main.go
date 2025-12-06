package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := ReadInputLines("day-4/input.txt")
	accessiblePaperCount := 0

	for rowIndex, line := range lines {
		for colIndex, char := range line {
			if char == '@' {
				adjacentPaperCount := countAdjacentPaper(lines, rowIndex, colIndex)
				if adjacentPaperCount < 4 {
					accessiblePaperCount++
				}
			}
		}
	}

	fmt.Println(accessiblePaperCount)
}

func countAdjacentPaper(lines []string, row, col int) int {
	rowStart := max(row-1, 0)
	rowEnd := min(row+1, len(lines)-1)
	colStart := max(col-1, 0)
	colEnd := min(col+1, len(lines[0])-1)
	adjacentPaperCount := 0

	for r := rowStart; r <= rowEnd; r++ {
		for c := colStart; c <= colEnd; c++ {
			if (r != row || c != col) && lines[r][c] == '@' {
				adjacentPaperCount++
			}
		}
	}

	return adjacentPaperCount
}

func ReadInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
