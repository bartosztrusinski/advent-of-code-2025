package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

func main() {
	sum := 0

	util.ScanInput("day-2/input.txt", func(input string) {
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
