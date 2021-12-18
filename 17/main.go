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

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	ocean := util.Grid{}


	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.Replace(line, ",", "", -1), " ")

		xParts := strings.Split(parts[2][2:], "..")
		yParts := strings.Split(parts[3][2:], "..")

		minX := util.Atoi(xParts[0])
		maxX := util.Atoi(xParts[1])
		minY := util.Atoi(yParts[0])
		maxY := util.Atoi(yParts[1])

		// Add Trench
		for i := minY; i <= maxY; i++ {
			for j := minX; j <= maxX; j++ {
				ocean.SetValue(j,i, "T")
			}
		}
	}

	// Set sub
	ocean.SetValue(0,0, "S")

	// TODO compute best velocity

	// Fire
	at := ocean.GetCoordinate(0,0)
	xVel := 6
	yVel := 9

	steps := 20

	// TODO know when to stop
	for i := 0;i < steps; i++ {
		at,xVel,yVel = updateProbe(at, xVel, yVel)
		ocean.SetCoordinate(at, at.Value)
	}

	// TODO // this will be the max height value
	//fmt.Println(ocean.GetMaxY())


	ocean.FlipVertically()
	ocean.FillGrid(".")
	ocean.PrintGrid(0)

	fmt.Printf("Part 1: %d\n", 0)
	fmt.Printf("Part 2: %d\n", 0)
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
