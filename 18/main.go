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

	pairStr := "[1,[[3,[2,[1,[7,3]]]],[6,1]]]" // TODO THIS IS NOT WORKING
	// need to better handle getting left and right.  I think the better way
	// is to traverse like the print function which is seeing nodes in order
	// then when we match getting the previous, and next.

	pair := parsePair(pairStr)

	fmt.Println(printPair(pair))
	fmt.Println(explode(pair))
	fmt.Println(printPair(pair))

	fmt.Printf("Part 1: %d\n", 0)
	fmt.Printf("Part 2: %d\n", 0)
}

func getDepth(current int, pair *Pair) (depth int) {

	xDepth := current
	if pair.xPair != nil {
		xDepth = getDepth(current + 1, pair.xPair)
	}

	yDepth := current
	if pair.yPair != nil {
		yDepth = getDepth(current + 1, pair.yPair)
	}

	max := xDepth

	if yDepth > max {
		max = yDepth
	}

	return max
}

func getFirstDepth4(current int, pair *Pair) (depth int, fpair *Pair) {

	fpair = pair
    xDepth := current
    if pair.xPair != nil {
        xDepth,fpair = getFirstDepth4(current + 1, pair.xPair)

		if xDepth >= 4 {
			return xDepth,fpair
		}
    }

    yDepth := current
    if pair.yPair != nil {
        yDepth,fpair = getFirstDepth4(current + 1, pair.yPair)

		if xDepth >= 4 {
			return xDepth,fpair
		}
    }

    max := yDepth

    if xDepth > max {
        max = xDepth
    }

    return max,fpair
}

func explode(pair *Pair) (bool, *Pair) {

	if getDepth(0, pair) < 4 {
		return false, pair
	}

	_,exploded := getFirstDepth4(0,pair)

	xPlode := exploded.xLiteral
	yPlode := exploded.yLiteral

	leftAt := exploded.parent
	rightAt := exploded.parent

	// set left value
	for leftAt != nil {
		if leftAt.xPair != nil {
			leftAt = leftAt.parent
		} else {
			leftAt.xLiteral += xPlode
			leftAt.xPair = nil
			break
		}
	}

	// set right value
	for rightAt != nil {
		if rightAt.yPair != nil {
			fmt.Println(rightAt)
			rightAt = rightAt.parent
		} else {
			rightAt.yLiteral += yPlode
			rightAt.yPair = nil
			break
		}
	}

	// replace with 0
	parent := exploded.parent

	if parent.xPair == exploded {
		parent.xPair = nil
	}

	if parent.yPair == exploded {
		parent.yPair = nil
	}



	return true, pair
}

func addPairs(a *Pair, b *Pair) *Pair {
	var pair *Pair
	new := Pair{}
	pair = &new
	pair.xPair = a
	pair.yPair = b
	return pair
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
