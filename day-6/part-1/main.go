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
	numbers, _ := parseToIntegers(lines[:len(lines)-1])
	sum := sumOperations(numbers, operands)
	fmt.Println(sum)
}

func sumOperations(numbers [][]int, operands []string) int {
	sum := 0
	rowSum := 0
	for col := range numbers[0] {
		isAddition := operands[col] == "+"
		if isAddition {
			rowSum = 0
		} else {
			rowSum = 1
		}
		for row := range numbers {
			if isAddition {
				rowSum += numbers[row][col]
			} else {
				rowSum *= numbers[row][col]
			}
		}
		sum += rowSum
	}
	return sum
}

func parseToIntegers(stringSlice []string) ([][]int, error) {
	intSlice := [][]int{}
	for _, line := range stringSlice {
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row[i] = value
		}
		intSlice = append(intSlice, row)
	}
	return intSlice, nil
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
