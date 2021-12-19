package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	//"stevee2112/aoc-2021/util"
	"strconv"
)

type Pair struct {
	xPair *Pair
	xLiteral int
	yPair *Pair
	yLiteral int
	parent *Pair
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		//_ := scanner.Text()
	}

	pairStr := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"

	pair := parsePair(pairStr)
	fmt.Println(pairStr)
	fmt.Println(printPair(pair))

	fmt.Printf("Part 1: %d\n", 0)
	fmt.Printf("Part 2: %d\n", 0)
}

func parsePair(p string) *Pair {
	var pair *Pair
	var at *Pair
	left := true

	for _,char := range strings.Split(p, "") {
		if char == "[" { // make a new Pair
			new := Pair{}
			if at != nil {
				if left {
					at.xPair = &new
				} else {
					at.yPair = &new
				}

				new.parent = at
			} else {
				pair = &new
			}

			left = true
			at = &new
		}

		if char == "]" {
			at = at.parent
		}

		if intChar, err := strconv.Atoi(char); err == nil { // literal
			if left {
				at.xLiteral = intChar
			} else {
				at.yLiteral = intChar
			}
		}

		if char == "," {
			left = false
		}
	}

	return pair
}

func printPair(p *Pair) string {

	output := "["

	if p.xPair == nil {
		output += strconv.Itoa(p.xLiteral)
	} else {
		output += printPair(p.xPair)
	}

	output += ","

	if p.yPair == nil {
		output += strconv.Itoa(p.yLiteral)
	} else {
		output += printPair(p.yPair)
	}

	output += "]"

	return output
}
