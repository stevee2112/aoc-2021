package main

import (
	"bufio"
	"fmt"
	"stevee2112/aoc-2021/util"
	"os"
	"path"
	"runtime"
	"strings"
	// "strconv"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	area1 := util.Grid{}
	area2 := util.Grid{}
	next := false

	y := 0
	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			next = true
			y = 0
			continue
		}

		if !next {
			x := 0
			for _,val := range strings.Split(line,"") {
				area1.SetValue(x, y, val)
				x++
			}
			y++
		} else {
			x := 0
			for _,val := range strings.Split(line,"") {
				area2.SetValue(x, y, val)
				x++
			}
			y++
		}
	}

	area1.PrintGrid(0)
	area2.PrintGrid(0)

	beacons1 := []Point{}

	scan1 := util.Coordinate{}
	area1.Traverse(func(coor util.Coordinate) bool {
		if coor.Value.(string) == "S" {
			scan1 = coor
			return false
		}

		return true
	})

	fmt.Println("here", scan1)

	area1.Traverse(func(coor util.Coordinate) bool {
		if coor.Value.(string) == "B" {
			fmt.Println(coor)
			beacons1 = append(beacons1, Point{coor.X - scan1.X, coor.Y - scan1.Y})
		}

		return true
	})

	fmt.Println("from grid", beacons1)

	// TODO be able to pull beacons from grid for easy testing

	beacons1 = []Point{
		Point{-1,1},
		Point{3,1},
		Point{2,3},
		Point{8,1},
	}

	fmt.Println("actual", beacons1)

	beacons2 := []Point{
		Point{0,-3},
		Point{6,1},
		Point{10,1},
		Point{9,3},
	}

	transformer1 := Point{}
	transformer2 := Point{}

	beacons1,transformer1 = getBounds(beacons1)
	beacons2,transformer2 = getBounds(beacons2)


	fmt.Println(beacons1)
	fmt.Println(beacons2)

	isMatch,relativeTransform := checkIfMatch(beacons1, beacons2)

	if isMatch {
		sxat := (-transformer2.X) + relativeTransform.X
		syat := (-transformer2.Y) + relativeTransform.Y

		fmt.Println("match found")
		fmt.Println("scanner 1 at 0,0")
		fmt.Println("scanner 2 at", Point{transformer1.X + sxat, transformer1.Y + syat})
	}

}

func checkIfMatch(a []Point, b []Point) (bool, Point) {

	matchesNeeded := 3

	axMax := a[0].X
	ayMax := a[0].Y
	bxMax := b[0].X
	byMax := b[0].Y
 	axMap := map[int]int{}
 	ayMap := map[int]int{}
 	bxMap := map[int]int{}
 	byMap := map[int]int{}


	for _,beacon := range a {
		if beacon.X > axMax {
			axMax = beacon.X
		}

		if beacon.Y > ayMax {
			ayMax = beacon.Y
		}

		axMap[beacon.X]++
		ayMap[beacon.Y]++
	}

	for _,beacon := range b {
		if beacon.X > bxMax {
			bxMax = beacon.X
		}

		if beacon.Y > byMax {
			byMax = beacon.Y
		}

		bxMap[beacon.X]++
		byMap[beacon.Y]++
	}

	// fmt.Println(ayMap, byMap)
	// fmt.Println(getMatchCount(ayMap, byMap))
	// fmt.Println(axMap, bxMap)
	// fmt.Println(getMatchCount(axMap, bxMap))

	// move box b up (y+) and check
	for i := 0;i <= byMax;i++ {
		bmoveboxy, currentbyMap := moveBox(b, "y", i)

		// if have min amount of match for y
		if getMatchCount(ayMap,currentbyMap) >= matchesNeeded {
			// move box right (x+) and check
			for j := 0;j <= bxMax;j++ {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					if getExactMatchCount(a, currentbbox) >= matchesNeeded {
						return true, Point{j,i}
					}
				}
			}
			// move box left and check
			for j := bxMax;j >= 0;j-- {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", axMax - j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					if getExactMatchCount(a, currentbbox) >= matchesNeeded {
						return true, Point{-j,i}
					}
				}
			}
		}
	}

	// move box b  (y-) and check
	for i := 0;i <= byMax;i++ {
		bmoveboxy, currentbyMap := moveBox(b, "y", -i)

		// if have min amount of match for y
		if getMatchCount(ayMap,currentbyMap) >= matchesNeeded {
			// move box right (x+) and check
			for j := 0;j <= bxMax;j++ {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					if getExactMatchCount(a, currentbbox) >= matchesNeeded {
						return true, Point{j,-i}
					}
				}
			}

			// move box left and check
			for j := 0;j <= bxMax;j++ {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", -j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					if getExactMatchCount(a, currentbbox) >= matchesNeeded {
						return true, Point{-j,-i}
					}
				}
			}
		}
	}

	// move box b down and check
	// if have min amount of match for y
		// move box right and check
		// move box left and check


	return false, Point{}
}

func getExactMatchCount(a []Point, b []Point) int {

	count := 0
	apointMap := map[string]bool{}
	bpointMap := map[string]bool{}

	for _,point := range a {
		apointMap[point.String()] = true
	}

	for _,point := range b {
		bpointMap[point.String()] = true
	}

	for point,_:= range apointMap {
		if _, exists := bpointMap[point]; exists {
			count++
		}
	}

	return count
}

func moveBox(box []Point, axis string, step int) ([]Point, map[int]int) {
	moved := []Point{}
	counts := map[int]int{}

	for _,point := range box {

		if axis == "x" {
			moved = append(moved, Point{point.X + step, point.Y})
			counts[point.X + step]++
		}

		if axis == "y" {
			moved = append(moved, Point{point.X, point.Y + step})
			counts[point.Y + step]++
		}
	}

	return moved,counts
}

func getMatchCount(a map[int]int, b map[int]int) int {

	matchCount := 0
	checkedMap := map[int]bool{}

	for point, count := range a {
		if _,checked := checkedMap[point]; !checked {
			checkedMap[point] = true
			matchCount += util.Min([]int{count, b[point]})
		}
	}

	for point, count := range b {
		if _,checked := checkedMap[point]; !checked {
			checkedMap[point] = true
			matchCount += util.Min([]int{count, a[point]})
		}
	}

	return matchCount
}

func getBounds(beacons []Point) ([]Point, Point) {

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

	return normalized, Point{xMin, yMin}
}
