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
		newDialPosition, pointedAtZeroCount := rotateDial(dialPosition, direction, distance)
		dialPosition = newDialPosition
		password += pointedAtZeroCount
	})

	fmt.Println(password)
}

func rotateDial(currentPosition int, direction rune, distance int) (int, int) {
	pointedAtZeroCount := distance / 100
	relevantDistance := distance % 100

	if direction == 'R' {
		newPosition := (currentPosition + relevantDistance) % 100
		if currentPosition+relevantDistance >= 100 {
			pointedAtZeroCount++
		}
		return newPosition, pointedAtZeroCount
	}

	newPosition := (currentPosition - relevantDistance + 100) % 100
	if currentPosition <= relevantDistance && currentPosition != 0 {
		pointedAtZeroCount++
	}

	return newPosition, pointedAtZeroCount
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
