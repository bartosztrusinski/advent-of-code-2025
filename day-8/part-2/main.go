package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLines("day-8/input.txt")
	positions, _ := linesToPositions(lines)
	_, lastConnection := connectShortestBoxes(positions)
	result := positions[lastConnection[0]][0] * positions[lastConnection[1]][0]
	fmt.Println(result)
}

func connectShortestBoxes(positions [][3]int) ([][]int, [2]int) {
	circuits := make([][]int, 0, len(positions))
	usedPairs := make(map[[2]int]bool)
	lastConnectionBoxIndices := [2]int{-1, -1}
	for len(circuits) != 1 || len(circuits[0]) != len(positions) {
		closestBoxDistance := math.MaxFloat64
		closestBoxIndices := [2]int{-1, -1}
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				pair := [2]int{i, j}
				if usedPairs[pair] {
					continue
				}
				distance := calculateDistance(positions[i], positions[j])
				if distance < closestBoxDistance {
					closestBoxDistance = distance
					closestBoxIndices = pair
				}
			}
		}
		usedPairs[closestBoxIndices] = true
		lastConnectionBoxIndices = closestBoxIndices
		circuitIndices := getCircuitIndices(circuits, closestBoxIndices)
		switch {
		case circuitIndices[0] != -1 && circuitIndices[0] == circuitIndices[1]:
			continue
		case circuitIndices[0] == -1 && circuitIndices[1] == -1:
			circuits = append(circuits, closestBoxIndices[:])
		case circuitIndices[0] != -1 && circuitIndices[1] == -1:
			circuits[circuitIndices[0]] = append(circuits[circuitIndices[0]], closestBoxIndices[1])
		case circuitIndices[0] == -1 && circuitIndices[1] != -1:
			circuits[circuitIndices[1]] = append(circuits[circuitIndices[1]], closestBoxIndices[0])
		case circuitIndices[0] != circuitIndices[1]:
			circuits[circuitIndices[0]] = append(circuits[circuitIndices[0]], circuits[circuitIndices[1]]...)
			circuits = append(circuits[:circuitIndices[1]], circuits[circuitIndices[1]+1:]...)
		}
	}
	return circuits, lastConnectionBoxIndices
}

func getCircuitIndices(circuits [][]int, boxIndices [2]int) [2]int {
	circuitIndices := [2]int{-1, -1}
	for index, circuit := range circuits {
		if slices.Contains(circuit, boxIndices[0]) {
			circuitIndices[0] = index
		}
		if slices.Contains(circuit, boxIndices[1]) {
			circuitIndices[1] = index
		}
	}
	return circuitIndices
}

func calculateDistance(p, q [3]int) float64 {
	sum := 0.0
	for i := range 3 {
		sum += math.Pow(float64(p[i]-q[i]), 2)
	}

	return math.Sqrt(sum)
}

func linesToPositions(lines []string) ([][3]int, error) {
	positions := make([][3]int, 0, len(lines))
	for _, line := range lines {
		coords := strings.Split(line, ",")
		position := [3]int{}
		for i, coord := range coords {
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
