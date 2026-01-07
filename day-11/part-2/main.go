package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInputLines("day-11/input.txt")
	graphMap := formatInput(lines)
	memo := make(map[string]int)
	visited := make(map[string]bool)
	pathCount := countOutputPaths(graphMap, "svr", "out", visited, memo)
	fmt.Println(pathCount)
}

func countOutputPaths(deviceMap map[string][]string, current string, target string, visited map[string]bool, memo map[string]int) int {
	if visited[current] {
		return 0
	}

	state := fmt.Sprintf("%v:%v:%v", current, visited["dac"], visited["fft"])
	if val, exists := memo[state]; exists {
		return val
	}

	visited[current] = true
	defer func() { visited[current] = false }()

	if current == target {
		if visited["dac"] && visited["fft"] {
			return 1
		}
		return 0
	}

	totalPaths := 0
	for _, neighbor := range deviceMap[current] {
		if !visited[neighbor] {
			totalPaths += countOutputPaths(deviceMap, neighbor, target, visited, memo)
		}
	}

	memo[state] = totalPaths
	return totalPaths
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
