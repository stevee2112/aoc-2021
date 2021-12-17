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

	riskMap := util.Grid{}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _,val := range strings.Split(line,"") {
			riskMap.SetValue(x, y, util.Atoi(val))
			x++
		}
		y++
	}

	frontier := riskMap.Frontier(
		riskMap.GetCoordinate(0,0),
		riskMap.GetCoordinate(riskMap.GetMaxX(), riskMap.GetMaxY()),
		func (at util.Coordinate, parent util.Coordinate, frontier util.Grid) (bool, interface{}) {
			sum :=  at.Value.(int) + parent.Value.(int)

			currentValue := frontier.GetCoordinate(at.X, at.Y)

			if currentValue.Value == nil || sum < currentValue.Value.(int) {
				return true, sum
			}

			return false, 0
		},
	)

	part1 := frontier.GetCoordinate(frontier.GetMaxX(), frontier.GetMaxY()).Value.(int)

	largeRiskMap := util.Grid{}

	for i := 0; i < 5; i++ {

		clone := riskMap.Clone()
		clone.Traverse(func(coor util.Coordinate) bool {
			newVal := (coor.Value.(int) + i)
			if newVal > 9 {
				newVal = newVal % 9
			}
			clone.SetValue(coor.X, coor.Y, newVal)
			return true
		});

		newRow := util.Grid{}
		for j := 0; j < 5; j++ {
			clone2 := clone.Clone()
			clone2.Traverse(func(coor util.Coordinate) bool {
				newVal := (coor.Value.(int) + j)
				if newVal > 9 {
					newVal = newVal % 9
				}
				clone2.SetValue(coor.X, coor.Y, newVal)
				return true
			});

			newRow = util.AppendHorizontal(newRow, clone2)

		}

		largeRiskMap = util.AppendVertical(largeRiskMap, newRow)
	}

	largeFrontier := largeRiskMap.Frontier(
		largeRiskMap.GetCoordinate(0,0),
		largeRiskMap.GetCoordinate(largeRiskMap.GetMaxX(), largeRiskMap.GetMaxY()),
		func (at util.Coordinate, parent util.Coordinate, frontier util.Grid) (bool, interface{}) {
			sum :=  at.Value.(int) + parent.Value.(int)

			currentValue := frontier.GetCoordinate(at.X, at.Y)
			if currentValue.Value == nil || sum < currentValue.Value.(int) {
				return true, sum
			}

			return false, 0
		},
	)

	part2 := largeFrontier.GetCoordinate(largeFrontier.GetMaxX(), largeFrontier.GetMaxY()).Value.(int)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
