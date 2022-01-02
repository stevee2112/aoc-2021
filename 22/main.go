package main

import (
	"fmt"
	"sort"
	//"strings"
	"stevee2112/aoc-2021/util"
)

type AxisRange struct {
	start int
	end int
}

type Range struct {
	xRange AxisRange
	yRange AxisRange
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

	ranges := removeRange([]Range{
		Range{
				AxisRange{0,3},
				AxisRange{0,3},
			},
		},
		Range{
			AxisRange{1,3},
			AxisRange{1,3},
		},
	);

	fmt.Println(len(ranges))
	fmt.Println(ranges)
	fmt.Println(sum(ranges))
}

func sum(ranges []Range) int {
	count := 0
	for _, aRange := range ranges {
		xDiff := util.Abs(aRange.xRange.end - aRange.xRange.start) + 1
		yDiff := util.Abs(aRange.yRange.end - aRange.yRange.start) + 1

		count += (xDiff * yDiff)
	}

	return count
}

func removeRange(ranges []Range, removeRange Range) []Range {
	newRanges := []Range{}

	xRanges := []AxisRange{}
	yRanges := []AxisRange{}

	for _,aRange := range ranges {
		xRanges = append(xRanges, aRange.xRange)
		yRanges = append(yRanges, aRange.yRange)
	}

	// add remove range
	xRanges = append(xRanges, removeRange.xRange)
	yRanges = append(yRanges, removeRange.yRange)

	uniqueXAxisRange := getUniqueAxisRanges(xRanges)
	uniqueYAxisRange := getUniqueAxisRanges(yRanges)

	for _, x := range uniqueXAxisRange{
		for _, y := range uniqueYAxisRange{

			// if remove range do not include it
			if checkIfOverlap(removeRange, Range{x, y}) {
				continue
			}

			include := false
			for _,original := range ranges {
				if checkIfOverlap(original, Range{x, y}) {
					include = true
					break
				}
			}

			if include {
				newRanges = append(newRanges, Range{x, y})
			}
		}
	}

	return newRanges
}

func mergeRange(ranges []Range) []Range {

	newRanges := []Range{}

	xRanges := []AxisRange{}
	yRanges := []AxisRange{}

	for _,aRange := range ranges {
		xRanges = append(xRanges, aRange.xRange)
		yRanges = append(yRanges, aRange.yRange)
	}

	uniqueXAxisRange := getUniqueAxisRanges(xRanges)
	uniqueYAxisRange := getUniqueAxisRanges(yRanges)

	for _, x := range uniqueXAxisRange{
		for _, y := range uniqueYAxisRange{

			include := false
			for _,original := range ranges {
				if checkIfOverlap(original, Range{x, y}) {
					include = true
					break
				}
			}

			if include {
				newRanges = append(newRanges, Range{x, y})
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
				newRanges = append(newRanges, AxisRange{previousVector.value, vector.value - 1})
			}
		}			

		if vector.label == "end" {
			if previousVector.label == "start" { // close current range
				newRanges = append(newRanges, AxisRange{previousVector.value, vector.value})
			} else {
				newRanges = append(newRanges, AxisRange{previousVector.value + 1, vector.value})
			}
			
		}

		previousVector = Vector{vector.value, vector.label}

	}

	return newRanges
}

func checkIfOverlap(a Range, b Range) bool {

    // If one rectangle is on left side of other
    if (a.xRange.start > b.xRange.end || b.xRange.start > a.xRange.end) {
		return false
	}

    // If one rectangle is above other
    if (a.yRange.start > b.yRange.end || b.yRange.start > a.yRange.end) {
		return false
	}

    return true;
}
