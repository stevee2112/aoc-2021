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

	floor := util.Grid{}

	y := 0

	for scanner.Scan() {

		line := scanner.Text()

		x := 0
		for _,val := range strings.Split(line,"") {
			floor.SetValue(x, y, val)
			x++
		}
		y++
	}

	at := 0

	for {
		movements := 0
		floor, movements = move(floor)

		at++

		if movements == 0 {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", at)
}

func move(floor util.Grid) (util.Grid, int) {

	rowMove  := 0
	floor, rowMove = moveRight(floor)

	colMove  := 0
	floor, colMove = moveDown(floor)


	return floor, (rowMove + colMove)
}

func moveRight(floor util.Grid) (util.Grid, int) {

	movements := 0
	rows := floor.GetRows()

	maxX := floor.GetMaxX()

	for _,row := range rows {
		var oZeroVal string
		for _,coor := range row {

			if coor.X == 0 {
				oZeroVal = coor.Value.(string)
			}

			if coor.Value.(string) == ">" {
				checkVal := ""
				nextX := coor.X + 1

				if nextX > maxX {
					nextX = 0
				}

				next := floor.GetCoordinate(nextX, coor.Y)

				if nextX == 0 {
					checkVal = oZeroVal
				} else {
					checkVal = next.Value.(string)
				}


				if checkVal == "." { // move
					movements++
					floor.SetValue(next.X, next.Y, ">")
					floor.SetValue(coor.X, coor.Y, ".")
				}
			} else {
				continue
			}
		}
	}

	return floor, movements
}

func moveDown(floor util.Grid) (util.Grid, int) {

	movements := 0
	cols := floor.GetCols()

	maxY := floor.GetMaxY()

	for _,col := range cols {
		var oZeroVal string
		for _,coor := range col {

			if coor.Y == 0 {
				oZeroVal = coor.Value.(string)
			}

			if coor.Value.(string) == "v" {
				checkVal := ""
				nextY := coor.Y + 1

				if nextY > maxY {
					nextY = 0
				}
				next := floor.GetCoordinate(coor.X, nextY)

				if nextY == 0 {
					checkVal = oZeroVal
				} else {
					checkVal = next.Value.(string)
				}

				if checkVal == "." { // move
					movements++
					floor.SetValue(next.X, next.Y, "v")
					floor.SetValue(coor.X, coor.Y, ".")
				}
			} else {
				continue
			}
		}
	}

	return floor, movements
}
