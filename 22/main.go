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
		"x=%d..%d,y=%d..%d,z=%d..%d",
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

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	ranges := []Range{}
	needPart1 := true

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

			if (intStart < -50 || intEnd > 50) && needPart1 {
				fmt.Println("Part 1:", sum(ranges))
				needPart1 = false
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
			ranges = mergeRange(ranges, newRange)
		}

		if action == "off" {
			ranges = removeRange(ranges, newRange)
		}
	}

	fmt.Println("Part 2:", sum(ranges))
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

	overlapRanges := []Range{}
	for _, aRange := range ranges {
		if checkIfOverlap(aRange, removeRange) {
			overlapRanges = append(overlapRanges, aRange)
		} else {
			newRanges = append(newRanges, aRange)
		}
	}

	for _,aRange := range overlapRanges {
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
				if checkIfInside(Range{x, y, z}, removeRange) {
					continue
				}

				include := false
				for _,original := range overlapRanges {
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

func mergeRange(ranges []Range, newRange Range) []Range {

	newRanges := []Range{}

	xRanges := []AxisRange{}
	yRanges := []AxisRange{}
	zRanges := []AxisRange{}

	overlapRanges := []Range{}
	for _, aRange := range ranges {
		if checkIfOverlap(aRange, newRange) {
			overlapRanges = append(overlapRanges, aRange)
		} else {
			newRanges = append(newRanges, aRange)
		}
	}

	// add new Range
	overlapRanges = append(overlapRanges, newRange)

	for _,aRange := range overlapRanges {
		xRanges = append(xRanges, aRange.xRange)
		yRanges = append(yRanges, aRange.yRange)
		zRanges = append(zRanges, aRange.zRange)
	}


	uniqueXAxisRange := getUniqueAxisRanges(xRanges)
	uniqueYAxisRange := getUniqueAxisRanges(yRanges)
	uniqueZAxisRange := getUniqueAxisRanges(zRanges)

	for _, x := range uniqueXAxisRange{
		for _, y := range uniqueYAxisRange{
			for _, z := range uniqueZAxisRange{

				rangeVal := Range{x, y, z}

				include := false

				for _,original := range overlapRanges {
					if checkIfOverlap(original, rangeVal) {
						include = true
						break
					}
				}

				if include {
					newRanges = append(newRanges, rangeVal)
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

	previousVector := vectors[0]
	currentRange := AxisRange{vectors[0].value, 0}

	for _, vector := range vectors {
		if vector.value == previousVector.value && vector.label == previousVector.label {
			continue
		}

		if vector.label == "start" {

			// close current range
			if (vector.value > currentRange.start) {
				currentRange.end = vector.value - 1
				newRanges = append(newRanges, currentRange)
			}

			newRange := AxisRange{vector.value, 0}
			currentRange = newRange
		}

		if vector.label == "end" {

			// close current range
			currentRange.end = vector.value
			newRanges = append(newRanges, currentRange)

			newRange := AxisRange{vector.value + 1, 0}
			currentRange = newRange
		}

		previousVector = vector
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

func checkIfInside(a Range, b Range) bool {

	if a.xRange.start < b.xRange.start {
		return false
	}

	if a.xRange.end > b.xRange.end {
		return false
	}

	if a.yRange.start < b.yRange.start {
		return false
	}

	if a.yRange.end > b.yRange.end {
		return false
	}

	if a.zRange.start < b.zRange.start {
		return false
	}

	if a.zRange.end > b.zRange.end {
		return false
	}

    return true;
}
