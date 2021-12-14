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

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	paper := util.Grid{}
	instructions := []string{}

	coordinates := true
	instructionsArea := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			coordinates = false
			instructionsArea = true
			continue
		}

		if coordinates {
			parts := strings.Split(line, ",")
			paper.SetValue(util.Atoi(parts[0]), util.Atoi(parts[1]), "#")
		}

		if instructionsArea {
			parts := strings.Split(line, " ")
			instructions = append(instructions, parts[2])
		}
	}

	part1Score := 0
	at := 0

	minX := 0
	maxX := paper.GetMaxX()
	minY := 0
	maxY := paper.GetMaxY()

	for _,instruction := range instructions {

		parts := strings.Split(instruction, "=")

		foldAt := util.Atoi(parts[1])

		if parts[0] == "y" {

			a := paper.Subset(minX,maxX,minY,(foldAt - 1))
			b := paper.Subset(minX,maxX,(foldAt + 1),maxY)
			b.FlipVertically()

			paper = util.MergeGrids(a, b)

			maxY = foldAt - 1
		}

		if parts[0] == "x" {

			a := paper.Subset(minX,(foldAt - 1),minY,maxY)
			b := paper.Subset((foldAt + 1),maxX,minY,maxY)
			b.FlipHorzontially()


			paper = util.MergeGrids(a, b)

			maxX = foldAt - 1
		}

		if at == 0 {
			paper.Traverse(func(coor util.Coordinate) bool {
				if coor.Value != nil && coor.Value.(string) == "#" {
					part1Score++
				}

				return true
			})
		}

		at++
	}


	paper.FillGrid(".")
	paper.PrintGrid(0)

	fmt.Printf("Part 1: %d\n", part1Score)
	fmt.Printf("Part 2: %d\n", 0)
}
