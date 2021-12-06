package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strconv"
	"strings"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	numbers := []string{}
	boards := []util.Grid{}
	haveNumbers := false
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		if !haveNumbers {
			numbers = strings.Split(line, ",")
			haveNumbers = true
			continue
		}

		if line == "" {
			// add grid
			boards = append(boards, util.Grid{})
			y = 0
			continue
		} else {
			parts := strings.Fields(line)

			for x, value := range parts {
				boards[len(boards) - 1].SetValue(x, y, strings.TrimSpace(value))
			}

			y++
		}
	}

	part1Answer := 0

	check:
	for _,number := range numbers {
		for _,board := range boards {
			board.Traverse(func(coor util.Coordinate) bool {
				if number == coor.Value {
					board.SetValue(coor.X, coor.Y, "X")
					return false
				}
				return true
			});
		}

		for _,board := range boards {
			if checkForWinner(board) {
				sum := 0
				board.Traverse(func(coor util.Coordinate) bool {
					if coor.Value != "X" {
						intVal,_ := strconv.Atoi(coor.Value.(string))
						sum += intVal
					}

					return true;
				});

				numberInt,_ := strconv.Atoi(number)
				part1Answer = sum * numberInt
				board.PrintGrid(3)
				break check
			}
		}
	}


	fmt.Printf("Part 1: %d\n", part1Answer)
	fmt.Printf("Part 2: %d\n", 0)
}

func checkForWinner(board util.Grid) bool {

	// row check here
	for _,row := range board.GetRows() {
		win := true
		for _,number := range row {
			if number.Value != "X" {
				win = false
				break
			}
		}

		if win {
			return true
		}
	}

	for _,col := range board.GetCols() {
		win := true
		for _,number := range col {
			if number.Value != "X" {
				win = false
				break
			}
		}

		if win {
			return true
		}
	}

	return false
}
