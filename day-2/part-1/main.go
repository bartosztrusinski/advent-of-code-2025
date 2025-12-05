package main

import (
	"bufio"
	"fmt"
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

func sumInvalidIds(firstId, lastId int) int {
	sum := 0

	for id := firstId; id <= lastId; id++ {
		stringId := strconv.Itoa(id)
		if len(stringId)%2 != 0 {
			continue
		}
		middleIndex := len(stringId) / 2
		firstHalf, secondHalf := stringId[:middleIndex], stringId[middleIndex:]
		if firstHalf == secondHalf {
			sum += id
		}
	}

	return sum
}
