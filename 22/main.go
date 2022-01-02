package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"path"
	"runtime"
	"strings"
	"stevee2112/aoc-2021/util"
)

type AxisRange struct {
	start int
	end int
}

type Range struct {
	xRange AxisRange
	yRange AxisRange
	zRange AxisRange
}

func (r Range) String() string {
	return fmt.Sprintf(
		"%d,%d,%d,%d,%d,%d",
		r.xRange.start,
		r.xRange.end,
		r.yRange.start,
		r.yRange.end,
		r.zRange.start,
		r.zRange.end,
	);
}

type Vector struct {
	value int
	label string
}

type Vectors []Vector

func (v Vectors) Len() int           { return len(v) }
func (v Vectors) Less(i, j int) bool {

	if v[i].value == v[j].value {
		if v[i].label == "start" {
			return true
		} else {
			return false
		}
	}

	return v[i].value < v[j].value
}
func (v Vectors) Swap(i, j int)      {
	v[i], v[j] = v[j], v[i]
}

func main() {

	// fmt.Println(getUniqueAxisRanges([]AxisRange{
	// 	AxisRange{1,1},
	// 	AxisRange{1,1},
	// }))

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	ranges := []Range{}

line:
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		action := parts[0]
		rangesStr := strings.Split(parts[1],",")

		newRange := Range{}

		for i, aRange := range rangesStr {
			parts = strings.Split(aRange[2:], "..")

			intStart := util.Atoi(parts[0])
			intEnd := util.Atoi(parts[1])

			if intStart < -50 {
				continue line
			}

			if intEnd > 50 {
				continue line
			}

			switch (i) {
			case 0:
				newRange.xRange = AxisRange{intStart, intEnd}
			case 1:
				newRange.yRange = AxisRange{intStart, intEnd}
			case 2:
				newRange.zRange = AxisRange{intStart, intEnd}
			}
		}

		if action == "on" {
			ranges = mergeRange(append(ranges, newRange))
		}

		if action == "off" {
			ranges = removeRange(ranges, newRange)
		}

		fmt.Println("after line", len(ranges))
		// for _, blah := range ranges {
		// 	fmt.Println(blah)
		// }
	}

	fmt.Println("Part 1:", sum(ranges))
}

func sum(ranges []Range) int {
	count := 0
	for _, aRange := range ranges {
		xDiff := util.Abs(aRange.xRange.end - aRange.xRange.start) + 1
		yDiff := util.Abs(aRange.yRange.end - aRange.yRange.start) + 1
		zDiff := util.Abs(aRange.zRange.end - aRange.zRange.start) + 1

		count += (xDiff * yDiff * zDiff)
	}

	return count
}

func removeRange(ranges []Range, removeRange Range) []Range {
	newRanges := []Range{}

	xRanges := []AxisRange{}
	yRanges := []AxisRange{}
	zRanges := []AxisRange{}

	for _,aRange := range ranges {
		xRanges = append(xRanges, aRange.xRange)
		yRanges = append(yRanges, aRange.yRange)
		zRanges = append(zRanges, aRange.zRange)
	}

	// add remove range
	xRanges = append(xRanges, removeRange.xRange)
	yRanges = append(yRanges, removeRange.yRange)
	zRanges = append(zRanges, removeRange.zRange)

	uniqueXAxisRange := getUniqueAxisRanges(xRanges)
	uniqueYAxisRange := getUniqueAxisRanges(yRanges)
	uniqueZAxisRange := getUniqueAxisRanges(zRanges)

	for _, x := range uniqueXAxisRange{
		for _, y := range uniqueYAxisRange{
			for _, z := range uniqueZAxisRange{

				// if remove range do not include it
				if checkIfOverlap(removeRange, Range{x, y, z}) {
					continue
				}

				include := false
				for _,original := range ranges {
					if checkIfOverlap(original, Range{x, y, z}) {
						include = true
						break
					}
				}

				if include {
					newRanges = append(newRanges, Range{x, y, z})
				}
			}
		}
	}

	return newRanges
}

func mergeRange(ranges []Range) []Range {

	newRanges := []Range{}

	xRanges := []AxisRange{}
	yRanges := []AxisRange{}
	zRanges := []AxisRange{}

	for _,aRange := range ranges {
		xRanges = append(xRanges, aRange.xRange)
		yRanges = append(yRanges, aRange.yRange)
		zRanges = append(zRanges, aRange.zRange)
	}

	uniqueXAxisRange := getUniqueAxisRanges(xRanges)
	uniqueYAxisRange := getUniqueAxisRanges(yRanges)
	uniqueZAxisRange := getUniqueAxisRanges(zRanges)

	rangeCache := map[string]bool{}

	for _, x := range uniqueXAxisRange{
		for _, y := range uniqueYAxisRange{
			for _, z := range uniqueZAxisRange{

				rangeVal := Range{x, y, z}

				if _,exists := rangeCache[rangeVal.String()]; !exists {
					include := false
					for _,original := range ranges {
						if checkIfOverlap(original, rangeVal) {
							include = true
							break
						}
					}

					if include {
						rangeCache[rangeVal.String()] = true
						newRanges = append(newRanges, rangeVal)
					}					
				}							
			}
		}
	}

	return newRanges
}

func getUniqueAxisRanges(ranges []AxisRange) []AxisRange {

	vectors := Vectors{}
	newRanges := []AxisRange{}

	for	_, aRange := range ranges {
		vectors = append(vectors, Vector{aRange.start, "start"})
		vectors = append(vectors, Vector{aRange.end, "end"})
	}

	sort.Sort(vectors)

	previousVector := Vector{vectors[0].value, "empty"}

	for _, vector := range vectors {
		if vector.label == "start" {
			if previousVector.label == "start" { // close current range

				// ignore values that are the same
				if vector.value == previousVector.value {
					continue
				}
				
				newRanges = append(newRanges, AxisRange{previousVector.value, vector.value - 1})
			}
		}			

		if vector.label == "end" {
			if previousVector.label == "start" { // close current range
				newRanges = append(newRanges, AxisRange{previousVector.value, vector.value})
			} else {

				// ignore values that are the same
				if vector.value == previousVector.value {
					continue
				}
				
				newRanges = append(newRanges, AxisRange{previousVector.value + 1, vector.value})
			}
			
		}

		previousVector = Vector{vector.value, vector.label}

	}

	return newRanges
}

func checkIfOverlap(a Range, b Range) bool {

    if (a.xRange.start > b.xRange.end || b.xRange.start > a.xRange.end) {
		return false
	}

    if (a.yRange.start > b.yRange.end || b.yRange.start > a.yRange.end) {
		return false
	}

    if (a.zRange.start > b.zRange.end || b.zRange.start > a.zRange.end) {
		return false
	}

    return true;
}
