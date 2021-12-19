package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"stevee2112/aoc-2021/util"
	//"strconv"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	ocean := util.Grid{}

	var minX, maxX, minY, maxY int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.Replace(line, ",", "", -1), " ")

		xParts := strings.Split(parts[2][2:], "..")
		yParts := strings.Split(parts[3][2:], "..")

		minX = util.Atoi(xParts[0])
		maxX = util.Atoi(xParts[1])
		minY = util.Atoi(yParts[0])
		maxY = util.Atoi(yParts[1])

		// Add Trench
		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				ocean.SetValue(j,i, "T")
			}
		}
	}

	// Set sub
	ocean.SetValue(0,0, "S")

	maxYhit := 0

	// try
	hits := 0
	iteration := 0
	for y := ocean.GetMinY();y <= util.Abs(ocean.GetMinY());y++ {
		for x := 0;x <= ocean.GetMaxX();x++ {
			try := ocean.Clone()
			at := try.GetCoordinate(0,0)
			xVel := x
			yVel := y

			hit := false

			for {
				at,xVel,yVel = updateProbe(at, xVel, yVel)
				try.SetCoordinate(at, at.Value)

				if at.X >= minX && at.X <= maxX && at.Y >= minY && at.Y <= maxY {
					hit = true
					break
				}

				if at.X > maxX || at.Y < minY {
					break
				}
			}

			if hit {
				hits++
				if try.GetMaxY() > maxYhit {
					maxYhit = try.GetMaxY()
				}

				// try.FlipVertically()
				// try.FillGrid(".")
				// try.PrintGrid(0)
			}

			iteration++
		}
	}

	fmt.Printf("Part 1: %d\n", maxYhit)
	fmt.Printf("Part 2: %d\n", hits)
}

func updateProbe(at util.Coordinate, xVel int, yVel int) (util.Coordinate, int, int) {
	xPos := at.X + xVel
	yPos := at.Y + yVel

	if xVel > 0 {
		xVel--
	}

	if xVel < 0 {
		xVel++
	}

	yVel--

	return util.Coordinate{xPos, yPos, "#"}, xVel, yVel
}
