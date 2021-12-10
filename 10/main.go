package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sort"
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

	part1Sum := 0
	part2Sums := []int{}
	completPoints := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	for _,line := range subsystem {
		corrupted, incomplete, badChar, completeString := checkLine(line)

		// Part 2
		if incomplete {
			sum := 0
			// compute score
			for _,char := range completeString {
				sum = (sum * 5) + completPoints[char]
			}

			part2Sums = append(part2Sums, sum)
			continue
		}

		// Part 1
		if corrupted {
			switch badChar {
			case ")":
				part1Sum += 3
			case "]":
				part1Sum += 57
			case "}":
				part1Sum += 1197
			case ">":
				part1Sum += 25137
			}
		}
	}

	sort.Ints(part2Sums)

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sums[len(part2Sums) / 2])
}

func checkLine(line []string) (corrupted bool, incomplete bool, char string, completeString []string) {

	endStartMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	startEndMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	if len(line) < 2 {
		return false, true, "", []string{}
	}

	startStack := []string{line[0]}
	for _,char := range line[1:] {

		if char == "(" || char == "[" || char == "{" || char == "<" { // chunk start
			startStack = append([]string{char}, startStack...)
		}

		if char == ")" || char == "]" || char == "}" || char == ">" { // chunk end
			startChar := endStartMap[char]

			if startChar != startStack[0] { // corrupted
				return true, false, char, []string{}
			}

			// pop stack
			startStack = startStack[1:]
		}
	}

	if len(startStack) > 0 { // incomplete
		completeString := []string{}

		for _,char := range startStack {
			completeString = append(completeString, startEndMap[char])
		}

		return false, true, "", completeString
	}

	return corrupted, incomplete, "", []string{}
}
