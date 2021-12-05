package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	//"strconv"
	"strings"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

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

	for _,board := range boards {
		board.PrintGrid(3)
	}

	fmt.Println(numbers)

	fmt.Printf("Part 1: %d\n", 0)
	fmt.Printf("Part 2: %d\n", 0)
}
