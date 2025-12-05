package main

import (
	"fmt"
	"strconv"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

func main() {
	dialPosition := 50
	password := 0

	util.ScanInput("day-1/input.txt", func(input string) {
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
