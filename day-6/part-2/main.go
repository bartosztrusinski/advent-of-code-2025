package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("day-6/input.txt")
	operands := strings.Fields(lines[len(lines)-1])
	parsedLines := parseLines(lines)
	sum := sumOperations(parsedLines, operands)
	fmt.Println(sum)

}

func sumOperations(lines [][]string, operands []string) int {
	sum := 0
	rowSum := 0
	for col := range lines[0] {
		isAddition := operands[col] == "+"
		if isAddition {
			rowSum = 0
		} else {
			rowSum = 1
		}
		for i := 0; i < len(lines[0][col]); i++ {
			rowStr := ""
			for row := range lines {
				char := string(lines[row][col][i])
				if char != " " {
					rowStr += char
				}
			}
			rowNumber, _ := strconv.Atoi(rowStr)
			if isAddition {
				rowSum += rowNumber
			} else {
				rowSum *= rowNumber
			}
		}
		sum += rowSum
	}
	return sum
}

func parseLines(lines []string) [][]string {
	operandRow := lines[len(lines)-1]
	columnStarts := operandPositions(operandRow)
	columnStarts = append(columnStarts, len(operandRow))

	matrix := make([][]string, 0, len(lines)-1)
	for _, line := range lines[:len(lines)-1] {
		matrix = append(matrix, sliceByColumns(line, columnStarts))
	}
	return matrix
}

func operandPositions(line string) []int {
	var positions []int
	for idx, ch := range line {
		if ch != ' ' {
			positions = append(positions, idx)
		}
	}
	return positions
}

func sliceByColumns(line string, columnBoundaries []int) []string {
	targetLen := columnBoundaries[len(columnBoundaries)-1]
	if len(line) < targetLen {
		line += strings.Repeat(" ", targetLen-len(line))
	}

	columns := make([]string, 0, len(columnBoundaries)-1)
	for i := 0; i < len(columnBoundaries)-1; i++ {
		start := columnBoundaries[i]
		end := columnBoundaries[i+1]
		if i < len(columnBoundaries)-2 && end-start > 0 {
			end--
		}
		width := end - start
		chunk := line[start:end]
		if len(chunk) < width {
			chunk += strings.Repeat(" ", width-len(chunk))
		}
		columns = append(columns, chunk)
	}
	return columns
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
