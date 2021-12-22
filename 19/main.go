package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "path"
	// "runtime"
	// "strings"
	// "strconv"
)

type Point struct {
	X int
	Y int
}

func main() {

	beacons1 := []Point{
		Point{3,3},
		Point{-5,2},
		Point{0,2},
		Point{4,1},
		Point{-5,-1},
	}

	beacons2 := []Point{
		Point{0,3},
		Point{-6,2},
		Point{-9,1},
		Point{0,1},
		Point{-5,0},
	}

	beacons1 = getBounds(beacons1)
	beacons2 = getBounds(beacons2)

	fmt.Println(checkIfMatch(beacons1, beacons2))
}

func checkIfMatch(a []Point, b []Point) (bool, Point) {
	axMax := a[0].X
	ayMax := a[0].Y
	bxMax := b[0].X
	byMax := b[0].Y

	for _,beacon := range a {
		if beacon.X > axMax {
			axMax = beacon.X
		}

		if beacon.Y > ayMax {
			ayMax = beacon.Y
		}
	}

	for _,beacon := range b {
		if beacon.X > bxMax {
			bxMax = beacon.X
		}

		if beacon.Y > byMax {
			byMax = beacon.Y
		}
	}

	// move box b up and check
	// if have min amount of match for y
		// move box right and check
		// move box left and check

	// move box b down and check
	// if have min amount of match for y
		// move box right and check
		// move box left and check


	return true, Point{}
}

func getBounds(beacons []Point) []Point {

	normalized := []Point{}

	xMin := beacons[0].X
	yMin := beacons[0].Y

	for _,beacon := range beacons {
		if beacon.X < xMin {
			xMin = beacon.X
		}

		if beacon.Y < yMin {
			yMin = beacon.Y
		}
	}

	for _,beacon := range beacons {
		normalized = append(normalized, Point{
			beacon.X - xMin,
			beacon.Y - yMin,
		})
	}

	return normalized
}
