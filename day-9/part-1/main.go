package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("day-9/input.txt")
	positions, _ := linesToPositions(lines)
	largestRectangle := findLargestRectangle(positions)
	fmt.Println(largestRectangle)
}

func findLargestRectangle(positions [][2]int) int {
	largestArea := 0
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			area := calculateArea(positions[i], positions[j])
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea
}

func calculateArea(cornerA, cornerB [2]int) int {
	width := math.Abs(float64(cornerA[0]-cornerB[0])) + 1
	height := math.Abs(float64(cornerA[1]-cornerB[1])) + 1
	return int(width * height)
}

func linesToPositions(lines []string) ([][2]int, error) {
	positions := make([][2]int, 0, len(lines))
	for _, line := range lines {
		position := [2]int{}
		for i, coord := range strings.Split(line, ",") {
			value, err := strconv.Atoi(coord)
			if err != nil {
				return nil, err
			}
			position[i] = value
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}
