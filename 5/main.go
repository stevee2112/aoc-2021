package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
)

type Vector struct {
	start util.Coordinate
	end   util.Coordinate
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	vectors := []Vector{}

	maxX := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "->", ",")

		parts := strings.Split(line, ",")

		// start X, end X
		if util.Atoi(parts[0]) > maxX {
			maxX = util.Atoi(parts[0])
		}

		if util.Atoi(parts[2]) > maxX {
			maxX = util.Atoi(parts[2])
		}

		// start Y, end Y
		if util.Atoi(parts[1]) > maxY {
			maxY = util.Atoi(parts[1])
		}

		if util.Atoi(parts[3]) > maxY {
			maxY = util.Atoi(parts[3])
		}

		vectors = append(vectors, Vector{
			start: util.Coordinate{util.Atoi(parts[0]), util.Atoi(parts[1]), nil},
			end:   util.Coordinate{util.Atoi(parts[2]), util.Atoi(parts[3]), nil},
		})
	}

	grid := util.MakeFullGrid(maxX, maxY, 0)
	//grid.PrintGrid(0)

	// Part 1
	for _,vector := range vectors {
		grid = applyVectorPart1(vector, grid)
	}

	//grid.PrintGrid(0)

	count := 0
	grid.Traverse(func(coor util.Coordinate) bool {
		intVal := coor.Value.(int)

		if intVal >= 2 {
			count++
		}
		return true
	});

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", 0)

}

func applyVectorPart1(vector Vector, grid util.Grid) util.Grid {

	// ignore non straight vectors
	if (vector.start.X != vector.end.X) && (vector.start.Y != vector.end.Y) {
		return grid
	}

	// Compute diff

	startXVal := vector.start.X
	startYVal := vector.start.Y
	xDiff := vector.end.X - vector.start.X

	if xDiff < 0 {
		startXVal = vector.end.X
		xDiff = vector.start.X - vector.end.X
	}

	yDiff := vector.end.Y - vector.start.Y

	if yDiff < 0 {
		startYVal = vector.end.Y
		yDiff = vector.start.Y - vector.end.Y
	}

	for i := xDiff; i >= 0;i-- {
		for j := yDiff; j >= 0;j-- {
			x := i + startXVal
			y := j + startYVal

			coor := grid.GetCoordinate(x, y)

			newValue := coor.Value.(int) + 1
			grid.SetValue(x, y, newValue)
		}
	}

	return grid
}
