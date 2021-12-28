package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
	"strconv"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	start := true
	transformMap := []string{}
	image := util.Grid{}

	y := 0
	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			start = false
			continue
		}

		if start {
			transformMap = strings.Split(line, "")
		} else {
			x := 0
			for _,val := range strings.Split(line,"") {
				image.SetValue(x, y, val)
				x++
			}
			y++
		}
	}

	infiniteChar := "."
	part1 := 0
	part2 := 0


	count := 50
	for i := 0;i < count;i++ {
		image = enhance(image, transformMap, infiniteChar)
		infiniteChar = getNextInfinite(infiniteChar, transformMap)

		// Part 1
		if i == 1 {
			image.Traverse(func(coor util.Coordinate) bool {
				if coor.Value.(string) == "#" {
					part1++
				}

				return true
			})
		}
	}

	image.Traverse(func(coor util.Coordinate) bool {
		if coor.Value.(string) == "#" {
			part2++
		}

		return true
	})

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func getNextInfinite(infiniteChar string, transformMap []string) string {
	infiniteStr := ""
	for i := 0;i < 9;i++ {
		if infiniteChar == "." {
			infiniteStr += "0"
		}

		if infiniteChar == "#" {
			infiniteStr += "1"
		}
	}

	i, _ := strconv.ParseInt(infiniteStr, 2, 64);

	return transformMap[i]

	return infiniteStr
}

func enhance(image util.Grid, transformMap []string, infiniteChar string) util.Grid {
	enhanced := util.Grid{}

	for i := image.GetMinX() - 2; i <= image.GetMaxX() + 2;i++ {
		for j := image.GetMinY() - 2; j <= image.GetMaxY() + 2;j++ {
			val := enhancePixel(image, i, j, transformMap, infiniteChar)
			enhanced.SetValue(i, j, val)
		}
	}

	return enhanced
}

func enhancePixel(g util.Grid, x int, y int, transformMap []string, infiniteChar string) string {
	binary := ""

	// Above left
	binary += valToIntStr(g.GetCoordinate(x - 1, y - 1).Value, infiniteChar)

	// Above
	binary += valToIntStr(g.GetCoordinate(x, y - 1).Value, infiniteChar)

	// Above right
	binary += valToIntStr(g.GetCoordinate(x + 1, y - 1).Value, infiniteChar)

	// Left
	binary += valToIntStr(g.GetCoordinate(x - 1, y).Value, infiniteChar)

	// Center
	binary += valToIntStr(g.GetCoordinate(x, y).Value, infiniteChar)

	// Right
	binary += valToIntStr(g.GetCoordinate(x + 1, y).Value, infiniteChar)

	// Below left
	binary += valToIntStr(g.GetCoordinate(x - 1, y + 1).Value, infiniteChar)

	// Below
	binary += valToIntStr(g.GetCoordinate(x, y + 1).Value, infiniteChar)

	// Below Right
	binary += valToIntStr(g.GetCoordinate(x + 1, y + 1).Value, infiniteChar)


	i, _ := strconv.ParseInt(binary, 2, 64);

	return transformMap[i]
}

func valToIntStr(val interface{}, infiniteChar string) string {

	strVal := ""

	if val == nil {
		strVal = infiniteChar
	} else {
		strVal = val.(string)
	}

	switch (strVal) {
	case ".":
		return "0"
	case "#":
		return "1"
	}

	return "0"
}
