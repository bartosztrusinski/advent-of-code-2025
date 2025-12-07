package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("day-5/input.txt")
	emptyLineIndex := slices.Index(lines, "")
	freshIngredientIdRanges := parseRanges(lines[:emptyLineIndex])
	availableIngredientIds := stringToIntSlice(lines[emptyLineIndex+1:])
	freshIngredientCount := 0

	for _, availableIngredientId := range availableIngredientIds {
		for _, freshIngredientIdRange := range freshIngredientIdRanges {
			if availableIngredientId >= freshIngredientIdRange[0] && availableIngredientId <= freshIngredientIdRange[1] {
				freshIngredientCount++
				break
			}
		}
	}

	fmt.Println(freshIngredientCount)
}

func parseRanges(ranges []string) [][2]int {
	parsedRange := [][2]int{}

	for _, r := range ranges {
		start, end := 0, 0
		fmt.Sscanf(r, "%d-%d", &start, &end)
		parsedRange = append(parsedRange, [2]int{start, end})
	}

	return parsedRange
}

func stringToIntSlice(stringSlice []string) []int {
	intSlice := make([]int, len(stringSlice))

	for index, str := range stringSlice {
		intSlice[index], _ = strconv.Atoi(str)
	}

	return intSlice
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
