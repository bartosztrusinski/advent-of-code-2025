package main

import (
	"fmt"
	"strings"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

func main() {
	stepsSum := 0
	util.ScanInput("day-10/input.txt", func(input string) {
		splitLights := strings.Split(input, "] ")
		indicatorLightsDiagram := splitLights[0][1:]
		indicatorLights := formatIndicatorLights(indicatorLightsDiagram)
		splitWiring := strings.Split(splitLights[1], " {")[0]
		wiringSchematics := readWiringSchematics(strings.Split(splitWiring, " "))
		currentLights := make([][]bool, 1)
		currentLights[0] = make([]bool, len(indicatorLights))

		stepsCount := 0
		for ; !areLightsMatching(currentLights, indicatorLights); stepsCount++ {
			currentLights = applyWirings(currentLights, wiringSchematics)
		}
		stepsSum += stepsCount
	})
	fmt.Println(stepsSum)
}

func applyWirings(currentLights [][]bool, wirings [][]int) [][]bool {
	newLights := make([][]bool, 0)
	for _, wiring := range wirings {
		for _, light := range currentLights {
			newLight := make([]bool, len(light))
			copy(newLight, light)
			for _, wiringIndex := range wiring {
				newLight[wiringIndex] = !light[wiringIndex]
			}
			newLights = append(newLights, newLight)
		}
	}
	return newLights
}

func areLightsMatching(currentLights [][]bool, targetLights []bool) bool {
	for _, light := range currentLights {
		match := true
		for index := range light {
			if light[index] != targetLights[index] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func formatIndicatorLights(diagram string) []bool {
	lights := make([]bool, len(diagram))
	for i, char := range diagram {
		if char == '#' {
			lights[i] = true
		} else {
			lights[i] = false
		}
	}
	return lights
}

func readWiringSchematics(schematics []string) [][]int {
	wiringSchematics := make([][]int, 0, len(schematics))
	for _, schematic := range schematics {
		wiringSchematics = append(wiringSchematics, formatWiringSchematic(schematic))
	}
	return wiringSchematics
}

func formatWiringSchematic(schematic string) []int {
	wiring := make([]int, 0, len(schematic)/2)
	for i := 1; i < len(schematic)-1; i += 2 {
		wiring = append(wiring, int(schematic[i]-'0'))
	}
	return wiring
}
