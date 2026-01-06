package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInputLines("day-11/input.txt")
	graphMap := formatInput(lines)
	pathCount := countOutputPathsIterative(graphMap, "you", "out")
	fmt.Println(pathCount)
}

func countOutputPathsIterative(deviceMap map[string][]string, start string, target string) int {
	pathCounts := make(map[string]int)
	pathCounts[start] = 1

	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		if current == target {
			continue
		}

		for _, neighbor := range deviceMap[current] {
			pathCounts[neighbor] += pathCounts[current]
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return pathCounts[target]
}

func readInputLines(path string) []string {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return strings.Split(string(fileContent), "\n")
}

func formatInput(lines []string) map[string][]string {
	result := make(map[string][]string)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			values := strings.Fields(strings.TrimSpace(parts[1]))
			result[key] = values
		}
	}
	return result
}
