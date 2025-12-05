package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const filePath = "day-2/input.txt"

func main() {
	sum := 0

	scanInput(func(input string) {
		for idRange := range strings.SplitSeq(input, ",") {
			idRangeSlice := strings.Split(idRange, "-")
			firstId, _ := strconv.Atoi(idRangeSlice[0])
			lastId, _ := strconv.Atoi(idRangeSlice[1])
			sum += sumInvalidIds(firstId, lastId)
		}
	})

	fmt.Println(sum)
}

func sumInvalidIds(firstId, lastId int) int {
	sum := 0

	for id := firstId; id <= lastId; id++ {
		if id < 11 {
			continue
		}

		stringId := strconv.Itoa(id)
		idLength := len(stringId)
		divisors := findDivisors(idLength)

		for _, partLength := range divisors {
			arePartsEqual := true
			firstPart := stringId[:partLength]

			for i := 1; i < idLength/partLength; i++ {
				startIndex := i * partLength
				endIndex := startIndex + partLength
				if stringId[startIndex:endIndex] != firstPart {
					arePartsEqual = false
					break
				}
			}

			if arePartsEqual {
				sum += id
				break
			}
		}
	}

	return sum
}

func findDivisors(num int) []int {
	divisors := []int{1}
	numSqrt := int(math.Sqrt(float64(num)))

	for i := 2; i <= numSqrt; i++ {
		if num%i == 0 {
			divisors = append(divisors, i, num/i)
		}
	}

	return divisors
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
