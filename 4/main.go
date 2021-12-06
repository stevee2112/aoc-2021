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
	boards := []*util.Grid{}
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
			boards = append(boards, &util.Grid{})
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

	part1Found := false
	part1Answer := 0
	part2Answer := 0
	winners := 0;

	for _,number := range numbers {

		for _,board := range boards {

			if board == nil {
				continue
			}

			board.Traverse(func(coor util.Coordinate) bool {
				if number == coor.Value {
					board.SetValue(coor.X, coor.Y, "X")
					return false
				}
				return true
			});
		}

		winnersAt := []int{}
		for i,board := range boards {

			if board == nil {
				continue
			}

			if checkForWinner(*board) {

				winnersAt = append(winnersAt, i)
				winners++
			}
		}

		for _,winnerAt := range winnersAt {

			if !part1Found {
				sum := 0
				boards[winnerAt].Traverse(func(coor util.Coordinate) bool {
					if coor.Value != "X" {
						intVal,_ := strconv.Atoi(coor.Value.(string))
						sum += intVal
					}

					return true;
				});

				numberInt,_ := strconv.Atoi(number)
				part1Answer = sum * numberInt
				part1Found = true
			}

			currentBoard := boards[winnerAt]

			// remove board
			boards[winnerAt] = nil

			if winners == len(boards) {
				sum := 0
				currentBoard.Traverse(func(coor util.Coordinate) bool {
					if coor.Value != "X" {
						intVal,_ := strconv.Atoi(coor.Value.(string))
						sum += intVal
					}

					return true;
				});

				numberInt,_ := strconv.Atoi(number)
				part2Answer = sum * numberInt
			}

		}
	}


	fmt.Printf("Part 1: %d\n", part1Answer)
	fmt.Printf("Part 2: %d\n", part2Answer)
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
