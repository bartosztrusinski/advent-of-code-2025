package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	lines := readInputLines("day-5/input.txt")
	emptyLineIndex := slices.Index(lines, "")
	freshIngredientIdRanges := parseRanges(lines[:emptyLineIndex])
	totalFreshIngredientCount := 0

	for {
		hasChanges := false

		for outerIndex, outerIdRange := range freshIngredientIdRanges {
			if outerIdRange[0] == -1 {
				continue
			}

			for innerIndex, innerIdRange := range freshIngredientIdRanges {
				if outerIndex == innerIndex || innerIdRange[0] == -1 {
					continue
				}

				if (innerIdRange[0] >= outerIdRange[0] && innerIdRange[0] <= outerIdRange[1]) || (outerIdRange[0] >= innerIdRange[0] && outerIdRange[0] <= innerIdRange[1]) {
					freshIngredientIdRanges[outerIndex] = [2]int{
						min(outerIdRange[0], innerIdRange[0]),
						max(outerIdRange[1], innerIdRange[1]),
					}
					freshIngredientIdRanges[innerIndex][0] = -1
					hasChanges = true
					break
				}
			}
		}

		if !hasChanges {
			break
		}
	}

	for _, idRange := range freshIngredientIdRanges {
		if idRange[0] == -1 {
			continue
		}
		totalFreshIngredientCount += (idRange[1] - idRange[0] + 1)
	}

	fmt.Println(totalFreshIngredientCount)
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

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
