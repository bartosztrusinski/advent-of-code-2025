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
