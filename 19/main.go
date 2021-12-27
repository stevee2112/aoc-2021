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
	Z int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanners := [][]Point{}

	at := 0
	scanners = append(scanners, []Point{})
	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "scanner") {
			continue
		}

		if line == "" {
			at++
			scanners = append(scanners, []Point{})
			continue
		}

		parts := strings.Split(line, ",")
		scanners[at] = append(scanners[at], Point{util.Atoi(parts[0]), util.Atoi(parts[1]), util.Atoi(parts[2])})
	}

	fmt.Println(scanners)

	for i:=0; i < 24; i++ {

		a := scanners[0]
		b := rotatePoints(scanners[1], i)

		transformer1 := Point{}
		transformer2 := Point{}

		boundry1,transformer1 := getBounds(a)
		boundry2,transformer2 := getBounds(b)

		isMatch,relativeTransform := checkIfMatch(boundry1, boundry2)

		if isMatch {
			sxat := (-transformer2.X) + relativeTransform.X
			syat := (-transformer2.Y) + relativeTransform.Y
			szat := (-transformer2.Z) + relativeTransform.Z

			fmt.Println("match found at rotation index", i)
			fmt.Println("scanner 1 at 0,0,0")
			fmt.Println("scanner 2 at", Point{transformer1.X + sxat, transformer1.Y + syat, transformer1.Z + szat})
			break
		}

	}
}

func checkIfMatch(a []Point, b []Point) (bool, Point) {

	matchesNeeded := 3

	axMax := a[0].X
	ayMax := a[0].Y
	azMax := a[0].Z
	bxMax := b[0].X
	byMax := b[0].Y
	bzMax := b[0].Z
 	axMap := map[int]int{}
 	ayMap := map[int]int{}
 	azMap := map[int]int{}
 	bxMap := map[int]int{}
 	byMap := map[int]int{}
 	bzMap := map[int]int{}


	for _,beacon := range a {
		if beacon.X > axMax {
			axMax = beacon.X
		}

		if beacon.Y > ayMax {
			ayMax = beacon.Y
		}

		if beacon.Z > azMax {
			azMax = beacon.Z
		}

		axMap[beacon.X]++
		ayMap[beacon.Y]++
		azMap[beacon.Z]++
	}

	for _,beacon := range b {
		if beacon.X > bxMax {
			bxMax = beacon.X
		}

		if beacon.Y > byMax {
			byMax = beacon.Y
		}

		if beacon.Z > bzMax {
			bzMax = beacon.Z
		}

		bxMap[beacon.X]++
		byMap[beacon.Y]++
		bzMap[beacon.Z]++
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

					// move box out (z+) and check
					for k := 0;k <= bzMax;k++ {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{j,i,k}
							}
						}
					}

					// move box in (z-) and check
					for k := bzMax;k >= 0;k-- {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{j,i,-k}
							}
						}
					}
				}
			}

			// move box left (x-) and check
			for j := bxMax;j >= 0;j-- {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", axMax - j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					// move box out (z+) and check
					for k := 0;k <= bzMax;k++ {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{-j,i,k}
							}
						}
					}

					// move box in (z-) and check
					for k := bzMax;k >= 0;k-- {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{-j,i,-k}
							}
						}
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

					// move box out (z+) and check
					for k := 0;k <= bzMax;k++ {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{j,-i,k}
							}
						}
					}

					// move box right (z-) and check
					for k := bzMax;k >= 0;k-- {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{j,-i,-k}
							}
						}
					}
				}
			}

			// move box left (-x) and check
			for j := 0;j <= bxMax;j++ {

				currentbbox, currentbxMap := moveBox(bmoveboxy, "x", -j)

				if getMatchCount(axMap,currentbxMap) >= matchesNeeded {

					// move box right (z+) and check
					for k := 0;k <= bzMax;k++ {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{-j,-i,k}
							}
						}
					}

					// move box right (z-) and check
					for k := bzMax;k >= 0;k-- {

						currentbbox, currentbzMap := moveBox(currentbbox, "z", k)

						if getMatchCount(azMap,currentbzMap) >= matchesNeeded {
							if getExactMatchCount(a, currentbbox) >= matchesNeeded {
								return true, Point{-j,-i,-k}
							}
						}
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
			moved = append(moved, Point{point.X + step, point.Y, point.Z})
			counts[point.X + step]++
		}

		if axis == "y" {
			moved = append(moved, Point{point.X, point.Y + step, point.Z})
			counts[point.Y + step]++
		}

		if axis == "z" {
			moved = append(moved, Point{point.X, point.Y, point.Z + step})
			counts[point.Z + step]++
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
	zMin := beacons[0].Z

	for _,beacon := range beacons {
		if beacon.X < xMin {
			xMin = beacon.X
		}

		if beacon.Y < yMin {
			yMin = beacon.Y
		}

		if beacon.Z < zMin {
			zMin = beacon.Z
		}
	}

	for _,beacon := range beacons {
		normalized = append(normalized, Point{
			beacon.X - xMin,
			beacon.Y - yMin,
			beacon.Z - zMin,
		})
	}

	return normalized, Point{xMin, yMin, zMin}
}

func rotatePoints(points []Point, stage int) []Point {
	rotated := []Point{}

	for _,point := range points {
		rotated = append(rotated, rotate(point, stage))
	}

	return rotated
}

func rotate(p Point, stage int) Point {

	rotated := Point{p.X, p.Y, p.Z}
	switch (stage) {
	case 0:
		rotated = p
	case 1: // -90x
		rotated.X = 1 * p.X
		rotated.Y = 1 * p.Z
		rotated.Z = -1 * p.Y
	case 2: // -180x
		rotated.X = 1 * p.X
		rotated.Y = -1 * p.Y
		rotated.Z = -1 * p.Z
	case 3: // -270x
		rotated.X = 1 * p.X
		rotated.Y = -1 * p.Z
		rotated.Z = 1 * p.Y
	case 4: // -180y
		rotated.X = -1 * p.X
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.Z
	case 5: // -180y, -90x
		rotated.X = -1 * p.X
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = 1 * current.Z
		rotated.Z = -1 * current.Y
	case 6: // -180y, -180x
		rotated.X = -1 * p.X
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Y
		rotated.Z = -1 * current.Z
	case 7: // -180y, -270x
		rotated.X = -1 * p.X
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Z
		rotated.Z = 1 * current.Y
	case 8: // -90z
		rotated.X = 1 * p.Y
		rotated.Y = -1 * p.X
		rotated.Z = 1 * p.Z
	case 9: // -90z, -90x
		rotated.X = 1 * p.Y
		rotated.Y = -1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = 1 * current.Z
		rotated.Z = -1 * current.Y
	case 10: // -90z, -180x
		rotated.X = 1 * p.Y
		rotated.Y = -1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Y
		rotated.Z = -1 * current.Z
	case 11: // -90z, -270x
		rotated.X = 1 * p.Y
		rotated.Y = -1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Z
		rotated.Z = 1 * current.Y
	case 12: // 90z
		rotated.X = -1 * p.Y
		rotated.Y = 1 * p.X
		rotated.Z = 1 * p.Z
	case 13: // 90z, -90x
		rotated.X = -1 * p.Y
		rotated.Y = 1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = 1 * current.Z
		rotated.Z = -1 * current.Y
	case 14: // 90z, -180x
		rotated.X = -1 * p.Y
		rotated.Y = 1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Y
		rotated.Z = -1 * current.Z
	case 15: // 90z, -270x
		rotated.X = -1 * p.Y
		rotated.Y = 1 * p.X
		rotated.Z = 1 * p.Z
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Z
		rotated.Z = 1 * current.Y
	case 16: // 90y
		rotated.X = 1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.X
	case 17: // 90y, -90x
		rotated.X = 1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = 1 * current.Z
		rotated.Z = -1 * current.Y
	case 18: // 90y, -180x
		rotated.X = 1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Y
		rotated.Z = -1 * current.Z
	case 19: // 90y, -270x
		rotated.X = 1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = -1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Z
		rotated.Z = 1 * current.Y
	case 20: // -90y
		rotated.X = -1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = 1 * p.X
	case 21: // -90y, -90x
		rotated.X = -1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = 1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = 1 * current.Z
		rotated.Z = -1 * current.Y
	case 22: // -90y, -180x
		rotated.X = -1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = 1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Y
		rotated.Z = -1 * current.Z
	case 23: // -90y, -270x
		rotated.X = -1 * p.Z
		rotated.Y = 1 * p.Y
		rotated.Z = 1 * p.X
		current := Point{rotated.X, rotated.Y, rotated.Z}
		rotated.X = 1 * current.X
		rotated.Y = -1 * current.Z
		rotated.Z = 1 * current.Y
	}

	return rotated
}
