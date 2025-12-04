package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const filePath = "day-1/input.txt"

func main() {
	dialPosition := 50
	password := 0

	scanInput(func(input string) {
		direction := rune(input[0])
		distance, _ := strconv.Atoi(input[1:])
		dialPosition = rotateDial(dialPosition, direction, distance)

		if dialPosition == 0 {
			password++
		}
	})

	fmt.Println(password)
}

func rotateDial(currentPosition int, direction rune, distance int) int {
	if direction == 'R' {
		return (currentPosition + distance) % 100
	}

	newPosition := currentPosition - (distance % 100)
	if newPosition < 0 {
		return newPosition + 100
	}

	return newPosition
}

func scanInput(callback func(string)) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
