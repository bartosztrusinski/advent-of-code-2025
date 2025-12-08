package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const connectionsToMake = 1000
const circuitsToCount = 3

func main() {
	lines := readInputLines("day-8/input.txt")
	positions, _ := linesToPositions(lines)
	circuits := connectShortestBoxes(positions, connectionsToMake)
	largestCircuits := getTopCircuitSizes(circuits, circuitsToCount)
	result := multiplyCircuitSizes(largestCircuits)
	fmt.Println(result)
}

func connectShortestBoxes(positions [][3]int, connectionsToMake int) [][]int {
	circuits := [][]int{}
	usedPairs := make(map[[2]int]bool)
	for range connectionsToMake {
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
		circuitIndices := getCircuitIndices(circuits, closestBoxIndices)
		if circuitIndices[0] == -1 && circuitIndices[1] == -1 {
			circuits = append(circuits, closestBoxIndices[:])
		} else if circuitIndices[0] != -1 && circuitIndices[1] != -1 {
			if circuitIndices[0] != circuitIndices[1] {
				circuits[circuitIndices[0]] = append(circuits[circuitIndices[0]], circuits[circuitIndices[1]]...)
				circuits = append(circuits[:circuitIndices[1]], circuits[circuitIndices[1]+1:]...)
			}
		} else if circuitIndices[0] != -1 {
			circuits[circuitIndices[0]] = append(circuits[circuitIndices[0]], closestBoxIndices[1])
		} else {
			circuits[circuitIndices[1]] = append(circuits[circuitIndices[1]], closestBoxIndices[0])
		}
	}
	return circuits
}

func multiplyCircuitSizes(circuitSizes []int) int {
	product := 1
	for _, size := range circuitSizes {
		product *= size
	}
	return product
}

func getTopCircuitSizes(circuits [][]int, count int) []int {
	topCircuitSizes := []int{}
	for range count {
		largestSize := 0
		largestIndex := -1
		for index, circuit := range circuits {
			if len(circuit) > largestSize {
				largestSize = len(circuit)
				largestIndex = index
			}
		}
		if largestIndex != -1 {
			topCircuitSizes = append(topCircuitSizes, largestSize)
			circuits = append(circuits[:largestIndex], circuits[largestIndex+1:]...)
		}
	}
	return topCircuitSizes
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
