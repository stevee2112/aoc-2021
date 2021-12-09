package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	//"strconv"
	"strings"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	heatMap := util.Grid{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _,val := range strings.Split(line,"") {
			heatMap.SetValue(x, y, util.Atoi(val))
			x++
		}
		y++
	}

	// Part 1
	sum := 0
	heatMap.Traverse(func(coor util.Coordinate) bool {
		adjacent := heatMap.GetAdjacent(coor)

		isLowest := true
		for _,a := range adjacent {
			if a.Value.(int) <= coor.Value.(int) {
				isLowest = false
				break
			}
		}

		if isLowest {
			sum += (1 + coor.Value.(int))
		}

		return true
	})

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", 0)
}
