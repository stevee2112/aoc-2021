package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	//"stevee2112/aoc-2021/util"
	"strings"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	subsystem := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		lineChars := strings.Split(line,"")

		subsystem = append(subsystem, lineChars)
	}

	// Part 1
	sum := 0
	for _,line := range subsystem {
		corrupted, incomplete, badChar := checkLine(line)

		if incomplete {
			continue
		}

		if corrupted {
			switch badChar {
			case ")":
				sum += 3
			case "]":
				sum += 57
			case "}":
				sum += 1197
			case ">":
				sum += 25137
			}
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", 0)
}

func checkLine(line []string) (corrupted bool, incomplete bool, char string) {

	endStartMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	if len(line) < 2 {
		return false, true, ""
	}

	startStack := []string{line[0]}
	for _,char := range line[1:] {

		if char == "(" || char == "[" || char == "{" || char == "<" { // chunk start
			startStack = append([]string{char}, startStack...)
		}

		if char == ")" || char == "]" || char == "}" || char == ">" { // chunk end
			startChar := endStartMap[char]

			if startChar != startStack[0] { // corrupted
				return true, false, char
			}

			// pop stack
			startStack = startStack[1:]
		}
	}

	if len(startStack) > 0 { // incomplete
		return false, true, ""
	}

	return corrupted, incomplete, ""
}
