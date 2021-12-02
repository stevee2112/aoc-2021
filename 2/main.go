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

	sub := util.NewDirectedGraph(nil)
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
		case "down":
			compassDir = util.South
		}

		for i := 0; i < value; i++ {
			sub.Move(compassDir)
		}
	}

	fmt.Printf("Part 1: %d\n", util.Abs(sub.At().X) * util.Abs(sub.At().Y))
	fmt.Printf("Part 2: %d\n", 0)
}
