package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strings"
	"strconv"
	"stevee2112/aoc-2021/util"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	subPart1 := util.NewDirectedGraph(nil)
	subPart2 := util.NewDirectedGraph(nil)
	aim := 0

	for scanner.Scan() {
		instruction := scanner.Text()

		parts := strings.Split(instruction, " ")
		direction := parts[0]
		value,_ := strconv.Atoi(parts[1])

		compassDir := util.South // initialize
		switch (direction) {
		case "forward":
			compassDir = util.East
		case "up":
			compassDir = util.North
			aim -= value
		case "down":
			compassDir = util.South
			aim += value
		}

		for i := 0; i < value; i++ {
			subPart1.Move(compassDir)
		}

		if direction == "forward" {
			for i := 0; i < value; i++ {
				subPart2.Move(util.East)
			}

			for i := 0; i < (aim * value); i++ {
				subPart2.Move(util.South)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", util.Abs(subPart1.At().X) * util.Abs(subPart1.At().Y))
	fmt.Printf("Part 2: %d\n", util.Abs(subPart2.At().X) * util.Abs(subPart2.At().Y))
}
