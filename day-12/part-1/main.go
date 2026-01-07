package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Shape struct {
	cells     [][]bool
	width     int
	height    int
	rotations []Shape
}

type Region struct {
	width  int
	height int
	counts []int
}

func main() {
	file, _ := os.Open("day-12/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	index := 0
	shapes := parseShapes(lines, &index)
	regions := parseRegions(lines, &index)
	count := 0

	for _, region := range regions {
		grid := make([][]bool, region.height)
		for i := range grid {
			grid[i] = make([]bool, region.width)
		}

		if canFit(grid, shapes, region) {
			count++
		}
	}

	fmt.Println(count)
}

func parseShapes(lines []string, idx *int) []Shape {
	shapes := []Shape{}

	for *idx < len(lines) {
		line := strings.TrimSpace(lines[*idx])
		if line == "" {
			(*idx)++
			continue
		}

		if !strings.Contains(line, ":") {
			break
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 || len(parts[0]) != 1 {
			break
		}

		(*idx)++

		var cells [][]bool
		for *idx < len(lines) {
			line := lines[*idx]
			if line == "" {
				(*idx)++
				break
			}
			if !strings.ContainsAny(line, "#.") {
				break
			}

			row := make([]bool, len(line))
			for i, ch := range line {
				row[i] = ch == '#'
			}
			cells = append(cells, row)
			(*idx)++
		}

		if len(cells) > 0 && len(cells[0]) > 0 {
			shape := Shape{
				cells:  cells,
				width:  len(cells[0]),
				height: len(cells),
			}
			shape.rotations = getRotations(shape)
			shapes = append(shapes, shape)
		}
	}

	return shapes
}

func parseRegions(lines []string, idx *int) []Region {
	regions := []Region{}

	for *idx < len(lines) {
		line := strings.TrimSpace(lines[*idx])
		if line == "" {
			(*idx)++
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			(*idx)++
			continue
		}

		dimParts := strings.Split(parts[0], "x")
		width, _ := strconv.Atoi(dimParts[0])
		height, _ := strconv.Atoi(dimParts[1])

		countStrings := strings.Fields(parts[1])
		counts := make([]int, len(countStrings))
		for i, s := range countStrings {
			counts[i], _ = strconv.Atoi(s)
		}

		regions = append(regions, Region{
			width:  width,
			height: height,
			counts: counts,
		})
		(*idx)++
	}

	return regions
}

func getRotations(shape Shape) []Shape {
	rotations := []Shape{shape}
	current := shape

	for range 3 {
		current = rotateShape(current)
		isNew := true
		for _, existing := range rotations {
			if shapesEqual(current, existing) {
				isNew = false
				break
			}
		}
		if isNew {
			rotations = append(rotations, current)
		}
	}

	return rotations
}

func shapesEqual(a, b Shape) bool {
	if a.width != b.width || a.height != b.height {
		return false
	}
	for r := 0; r < a.height; r++ {
		for c := 0; c < a.width; c++ {
			if a.cells[r][c] != b.cells[r][c] {
				return false
			}
		}
	}
	return true
}

func rotateShape(shape Shape) Shape {
	h := shape.height
	w := shape.width
	newCells := make([][]bool, w)
	for i := range newCells {
		newCells[i] = make([]bool, h)
	}

	for r := range h {
		for c := range w {
			newCells[c][h-1-r] = shape.cells[r][c]
		}
	}

	return Shape{
		cells:  newCells,
		width:  w,
		height: h,
	}
}

func canPlace(grid [][]bool, shape Shape, row, col int) bool {
	if row+shape.height > len(grid) || col+shape.width > len(grid[0]) {
		return false
	}

	for r := 0; r < shape.height; r++ {
		for c := 0; c < shape.width; c++ {
			if shape.cells[r][c] && grid[row+r][col+c] {
				return false
			}
		}
	}
	return true
}

func place(grid [][]bool, shape Shape, row, col int) {
	for r := 0; r < shape.height; r++ {
		for c := 0; c < shape.width; c++ {
			if shape.cells[r][c] {
				grid[row+r][col+c] = true
			}
		}
	}
}

func unPlace(grid [][]bool, shape Shape, row, col int) {
	for r := 0; r < shape.height; r++ {
		for c := 0; c < shape.width; c++ {
			if shape.cells[r][c] {
				grid[row+r][col+c] = false
			}
		}
	}
}

func getShapeSize(shape Shape) int {
	count := 0
	for r := 0; r < shape.height; r++ {
		for c := 0; c < shape.width; c++ {
			if shape.cells[r][c] {
				count++
			}
		}
	}
	return count
}

func canFit(grid [][]bool, shapes []Shape, region Region) bool {
	presents := []int{}
	totalArea := 0
	for shapeID, count := range region.counts {
		for range count {
			presents = append(presents, shapeID)
			totalArea += getShapeSize(shapes[shapeID])
		}
	}

	gridArea := region.width * region.height
	if totalArea > gridArea {
		return false
	}

	sort.Slice(presents, func(i, j int) bool {
		return getShapeSize(shapes[presents[i]]) > getShapeSize(shapes[presents[j]])
	})

	return tryPlace(grid, shapes, presents, 0)
}

func tryPlace(grid [][]bool, allShapes []Shape, presents []int, idx int) bool {
	if idx == len(presents) {
		return true
	}

	shapeID := presents[idx]
	rotations := allShapes[shapeID].rotations

	for row := range grid {
		for col := 0; col < len(grid[0]); col++ {
			for _, rotation := range rotations {
				if canPlace(grid, rotation, row, col) {
					place(grid, rotation, row, col)
					if tryPlace(grid, allShapes, presents, idx+1) {
						unPlace(grid, rotation, row, col)
						return true
					}
					unPlace(grid, rotation, row, col)
				}
			}
		}
	}
	return false
}
