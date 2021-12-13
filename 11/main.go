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

	levels := util.Grid{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _,val := range strings.Split(line,"") {
			levels.SetValue(x, y, util.Atoi(val))
			x++
		}
		y++
	}


	// Part 1
	flashSum := 0
	stepCount := 100
	levels.PrintGrid(2)

	for i := 0; i < stepCount; i++ {
		new, flashCount := step(levels)
		levels = new
		flashSum += flashCount
	}

	levels.PrintGrid(2)

	fmt.Printf("Part 1: %d\n", flashSum)
	fmt.Printf("Part 2: %d\n", 0)
}

func step(current util.Grid) (util.Grid, int) {

	flashCount := 0
	flashes := []util.Coordinate{}
	flashed := []util.Coordinate{}

	// increase each number by 1
	current.Traverse(func(coor util.Coordinate) bool {
		intVal := coor.Value.(int) + 1

		current.SetValue(coor.X, coor.Y, intVal)

		if intVal == 10 {
			flashes = append(flashes, coor)
		}
		return true
	});

	//for each 10 flash (tens may increase so need to traverse whole grid each time or keep track of WHERE things need to flase)
	// - increase surround (if less than 10)
	// - set this to 11, (11 means flashed)
	for len(flashes) > 0 {

		flasher := flashes[0]
		surrounding := current.GetSurrounding(flasher)

		for _,coor := range surrounding {
			intVal := coor.Value.(int)

			if intVal < 10 {
				intVal++
				current.SetValue(coor.X, coor.Y, intVal)

				if intVal == 10 {
					flashes = append(flashes, coor)
				}
			}

		}

		current.SetValue(flasher.X, flasher.Y, 11)
		flashed = append(flashed, flasher)
		flashes = flashes[1:]
	}

	// when no more 10s set each 11 to 0
	flashCount = len(flashed)
	for _,coor := range flashed {
		current.SetValue(coor.X, coor.Y, 0)
	}

	return current, flashCount
}
