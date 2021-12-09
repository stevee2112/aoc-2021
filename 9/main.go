package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
	"sort"
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
	lowPoints := []util.Coordinate{}
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
			lowPoints = append(lowPoints, coor)
			sum += (1 + coor.Value.(int))
		}

		return true
	})

	// Part2
	basins := []int{}
	for _,lowPoint := range lowPoints {
		basins = append(basins, getBasinSize(heatMap, lowPoint))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	part2Answer := basins[0] * basins[1] * basins[2]

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", part2Answer)
}

func getBasinSize(heatMap util.Grid, coor util.Coordinate) int {

	inBasin := getCoordinatesInBasin(heatMap, coor, map[string]bool{})
	return len(inBasin)
}

func getCoordinatesInBasin(heatMap util.Grid, coor util.Coordinate, alreadyVisited map[string]bool) []util.Coordinate {

	alreadyVisited[coor.String()] = true

	coordinates := []util.Coordinate{coor}

	for _,a := range heatMap.GetAdjacent(coor) {
		if a.Value == 9 {
			continue
		} else {
			if _,found := alreadyVisited[a.String()]; found {
				continue // skip if we cam from this coordinate
			}

			coordinates = append(coordinates, getCoordinatesInBasin(heatMap, a, alreadyVisited)...)
		}
	}

	return coordinates
}
