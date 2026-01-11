package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/bartosztrusinski/advent-of-code-2025/util"
)

func main() {
	totalPresses := 0
	util.ScanInput("day-10/input.txt", func(input string) {
		buttons, targets := parseLine(input)
		presses := solveMachineILP(buttons, targets)
		totalPresses += presses
	})
	fmt.Println(totalPresses)
}

func parseLine(line string) ([][]int, []int) {
	buttonRe := regexp.MustCompile(`\(([^)]+)\)`)
	buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)

	buttons := make([][]int, len(buttonMatches))
	for i, match := range buttonMatches {
		parts := strings.Split(match[1], ",")
		buttons[i] = make([]int, len(parts))
		for j, p := range parts {
			buttons[i][j], _ = strconv.Atoi(p)
		}
	}

	targetRe := regexp.MustCompile(`\{([^}]+)\}`)
	targetMatch := targetRe.FindStringSubmatch(line)
	targetParts := strings.Split(targetMatch[1], ",")
	targets := make([]int, len(targetParts))
	for i, p := range targetParts {
		targets[i], _ = strconv.Atoi(p)
	}

	return buttons, targets
}

func solveMachineILP(buttons [][]int, targets []int) int {
	numButtons := len(buttons)
	numCounters := len(targets)

	A := make([][]float64, numCounters)
	for i := range A {
		A[i] = make([]float64, numButtons)
	}
	for j, button := range buttons {
		for _, counter := range button {
			if counter < numCounters {
				A[counter][j] = 1
			}
		}
	}
	b := make([]float64, numCounters)
	for i, t := range targets {
		b[i] = float64(t)
	}

	bestSolution := math.MaxInt32

	var branchAndBound func(lowerBounds, upperBounds []float64, depth int)
	branchAndBound = func(lowerBounds, upperBounds []float64, depth int) {
		solution, objVal, feasible := solveLPWithBounds(A, b, numButtons, numCounters, lowerBounds, upperBounds)

		if !feasible || objVal >= float64(bestSolution)-0.5 {
			return
		}

		fracIdx := -1
		maxFrac := 0.001
		for j := range numButtons {
			frac := solution[j] - math.Floor(solution[j])
			if frac > 0.001 && frac < 0.999 {
				distToHalf := math.Abs(frac - 0.5)
				if 0.5-distToHalf > maxFrac {
					maxFrac = 0.5 - distToHalf
					fracIdx = j
				}
			}
		}

		if fracIdx == -1 {
			total := 0
			for _, v := range solution {
				total += int(math.Round(v))
			}
			if total < bestSolution {
				bestSolution = total
			}
			return
		}

		floorVal := math.Floor(solution[fracIdx])
		ceilVal := math.Ceil(solution[fracIdx])

		newUpper := make([]float64, numButtons)
		newLower := make([]float64, numButtons)
		copy(newUpper, upperBounds)
		copy(newLower, lowerBounds)

		newLower2 := make([]float64, numButtons)
		copy(newLower2, newLower)
		newLower2[fracIdx] = ceilVal
		branchAndBound(newLower2, newUpper, depth+1)

		newUpper2 := make([]float64, numButtons)
		copy(newUpper2, newUpper)
		newUpper2[fracIdx] = floorVal
		branchAndBound(newLower, newUpper2, depth+1)
	}

	lower := make([]float64, numButtons)
	upper := make([]float64, numButtons)
	maxTarget := 0
	for _, t := range targets {
		if t > maxTarget {
			maxTarget = t
		}
	}
	for j := range upper {
		upper[j] = float64(maxTarget)
	}

	branchAndBound(lower, upper, 0)

	return bestSolution
}

func solveLPWithBounds(A [][]float64, b []float64, numVars, numConstraints int, lower, upper []float64) ([]float64, float64, bool) {
	m := numConstraints
	n := numVars
	M := 1e12

	bAdj := make([]float64, m)
	for i := range m {
		bAdj[i] = b[i]
		for j := 0; j < n; j++ {
			bAdj[i] -= A[i][j] * lower[j]
		}
		if bAdj[i] < -1e-9 {
			return nil, 0, false
		}
	}

	numSlack := n
	numArt := m + n
	totalVars := n + numSlack + numArt
	totalCons := m + n

	cols := totalVars + 1
	tableau := make([][]float64, totalCons+1)
	for i := range tableau {
		tableau[i] = make([]float64, cols)
	}

	for i := range m {
		for j := range n {
			tableau[i][j] = A[i][j]
		}
		tableau[i][n+numSlack+i] = 1
		tableau[i][cols-1] = bAdj[i]
	}

	for j := range n {
		ub := upper[j] - lower[j]
		if ub < -1e-9 {
			return nil, 0, false
		}
		tableau[m+j][j] = 1
		tableau[m+j][n+j] = 1
		if ub < 1e9 {
			tableau[m+j][n+numSlack+m+j] = 1
			tableau[m+j][cols-1] = ub
		} else {
			tableau[m+j][cols-1] = 1e9
		}
	}

	for j := range n {
		tableau[totalCons][j] = 1
	}
	for j := n + numSlack; j < totalVars; j++ {
		tableau[totalCons][j] = M
	}

	basic := make([]int, totalCons)
	for i := range m {
		basic[i] = n + numSlack + i
		for j := range cols {
			tableau[totalCons][j] -= M * tableau[i][j]
		}
	}
	for i := range n {
		ub := upper[i] - lower[i]
		if ub < 1e9 {
			basic[m+i] = n + numSlack + m + i
			for j := range cols {
				tableau[totalCons][j] -= M * tableau[m+i][j]
			}
		} else {
			basic[m+i] = n + i
		}
	}

	simplex(tableau, basic, totalCons, totalVars, cols)

	for i := range totalCons {
		if basic[i] >= n+numSlack && tableau[i][cols-1] > 1e-6 {
			return nil, 0, false
		}
	}

	solution := make([]float64, numVars)
	for i := range totalCons {
		if basic[i] < n {
			solution[basic[i]] = tableau[i][cols-1]
		}
	}
	for j := range n {
		solution[j] += lower[j]
		if solution[j] < lower[j] {
			solution[j] = lower[j]
		}
		if solution[j] > upper[j] {
			solution[j] = upper[j]
		}
	}

	objVal := 0.0
	for _, v := range solution {
		objVal += v
	}

	return solution, objVal, true
}

func simplex(tableau [][]float64, basic []int, m, n, cols int) {
	for range 100000 {
		pivotCol := -1
		minVal := -1e-9
		for j := 0; j < n; j++ {
			if tableau[m][j] < minVal {
				minVal = tableau[m][j]
				pivotCol = j
			}
		}
		if pivotCol == -1 {
			break
		}

		pivotRow := -1
		minRatio := math.Inf(1)
		for i := range m {
			if tableau[i][pivotCol] > 1e-9 {
				ratio := tableau[i][cols-1] / tableau[i][pivotCol]
				if ratio >= -1e-9 && ratio < minRatio-1e-9 {
					minRatio = ratio
					pivotRow = i
				}
			}
		}
		if pivotRow == -1 {
			break
		}

		pivotVal := tableau[pivotRow][pivotCol]
		for j := 0; j < cols; j++ {
			tableau[pivotRow][j] /= pivotVal
		}
		for i := 0; i <= m; i++ {
			if i != pivotRow {
				factor := tableau[i][pivotCol]
				for j := range tableau[i] {
					tableau[i][j] -= factor * tableau[pivotRow][j]
				}
			}
		}
		basic[pivotRow] = pivotCol
	}
}
