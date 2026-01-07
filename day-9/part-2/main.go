package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Edge struct {
	a Point
	b Point
}

func main() {
	vertices, _ := readPositions("day-9/input.txt")
	edges := buildEdges(vertices)
	largestRectangle := findLargestRectangle(vertices, edges)
	fmt.Println(largestRectangle)
}

func findLargestRectangle(vertices []Point, edges []Edge) int {
	insideCache := make(map[Point]bool)
	largestArea := 0

	for i := range vertices {
		for j := i + 1; j < len(vertices); j++ {
			if vertices[i].x == vertices[j].x || vertices[i].y == vertices[j].y {
				continue
			}

			minX := min(vertices[i].x, vertices[j].x)
			maxX := max(vertices[i].x, vertices[j].x)
			minY := min(vertices[i].y, vertices[j].y)
			maxY := max(vertices[i].y, vertices[j].y)

			area := (maxX - minX + 1) * (maxY - minY + 1)
			if area <= largestArea {
				continue
			}

			cornerA := Point{x: vertices[i].x, y: vertices[j].y}
			cornerB := Point{x: vertices[j].x, y: vertices[i].y}

			if !pointInsidePolygon(vertices[i], edges, insideCache) ||
				!pointInsidePolygon(vertices[j], edges, insideCache) ||
				!pointInsidePolygon(cornerA, edges, insideCache) ||
				!pointInsidePolygon(cornerB, edges, insideCache) {
				continue
			}

			if polygonCutsRectangle(minX, maxX, minY, maxY, edges) {
				continue
			}

			largestArea = area
		}
	}

	return largestArea
}

func pointInsidePolygon(p Point, edges []Edge, cache map[Point]bool) bool {
	if cached, ok := cache[p]; ok {
		return cached
	}
	inside := pointInOrOnPolygon(p, edges)
	cache[p] = inside
	return inside
}

func pointInOrOnPolygon(p Point, edges []Edge) bool {
	inside := false
	for _, edge := range edges {
		if pointOnSegment(p, edge) {
			return true
		}
		if edge.a.y == edge.b.y {
			continue
		}
		if (edge.a.y > p.y) == (edge.b.y > p.y) {
			continue
		}
		if edge.a.x > p.x {
			inside = !inside
		}
	}
	return inside
}

func pointOnSegment(p Point, edge Edge) bool {
	if edge.a.x == edge.b.x {
		if p.x != edge.a.x {
			return false
		}
		minY := min(edge.a.y, edge.b.y)
		maxY := max(edge.a.y, edge.b.y)
		return p.y >= minY && p.y <= maxY
	}
	if edge.a.y == edge.b.y {
		if p.y != edge.a.y {
			return false
		}
		minX := min(edge.a.x, edge.b.x)
		maxX := max(edge.a.x, edge.b.x)
		return p.x >= minX && p.x <= maxX
	}
	return false
}

func polygonCutsRectangle(minX, maxX, minY, maxY int, edges []Edge) bool {
	for _, edge := range edges {
		if edge.a.x == edge.b.x {
			x := edge.a.x
			if x <= minX || x >= maxX {
				continue
			}

			segMinY := min(edge.a.y, edge.b.y)
			segMaxY := max(edge.a.y, edge.b.y)
			if max(segMinY, minY+1) <= min(segMaxY, maxY-1) {
				return true
			}

			continue
		}

		if edge.a.y == edge.b.y {
			y := edge.a.y
			if y <= minY || y >= maxY {
				continue
			}

			segMinX := min(edge.a.x, edge.b.x)
			segMaxX := max(edge.a.x, edge.b.x)
			if max(segMinX, minX+1) <= min(segMaxX, maxX-1) {
				return true
			}
		}
	}

	return false
}

func buildEdges(vertices []Point) []Edge {
	edges := make([]Edge, 0, len(vertices))
	for i := range vertices {
		next := (i + 1) % len(vertices)
		edges = append(edges, Edge{a: vertices[i], b: vertices[next]})
	}
	return edges
}

func readPositions(path string) ([]Point, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(fileContent)), "\n")
	positions := make([]Point, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}

		positions = append(positions, Point{x: x, y: y})
	}

	return positions, nil
}
