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

	// TODO THIS EXAMPLE IS WRONG SPLIT ORDER IS WRONG
	pairStr1 := "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]" 
	pair1 := parsePair(pairStr1)

	pairStr2 := "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]"
	pair2 := parsePair(pairStr2)

	fmt.Println(printPair(pair1))
	fmt.Println(printPair(pair2))

	pair := addPairs(pair1, pair2)

	fmt.Println("original\t", printPair(pair))
	for {
		exploded :=explode(pair)

		if exploded {
			fmt.Println("explode\t\t", printPair(pair))
			continue
		}

		splited :=split(pair)

		if splited {
			fmt.Println("split\t\t", printPair(pair))
			continue
		}

		break
	}

	fmt.Println("final\t\t", printPair(pair))

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

func split(pair *Pair) (bool) {

	var toSplit *Pair
	isLeft := true
	traverse(pair, func(at *Pair) bool {
		if at.xLiteral >= 10 {
			toSplit = at
			isLeft = true
			return false
		}

		if at.yLiteral >= 10 {
			toSplit = at
			isLeft = false
			return false
		}

		return true
	})

	if toSplit != nil {

		newPair := Pair{}

		if isLeft {
			newPair.xLiteral = toSplit.xLiteral / 2
			newPair.yLiteral = (toSplit.xLiteral / 2) + (toSplit.xLiteral % 2)
			newPair.parent = toSplit
			toSplit.xLiteral = 0
			toSplit.xPair = &newPair
		} else {
			newPair.xLiteral = toSplit.yLiteral / 2
			newPair.yLiteral = (toSplit.yLiteral / 2) + (toSplit.yLiteral % 2)
			newPair.parent = toSplit
			toSplit.yLiteral = 0
			toSplit.yPair = &newPair
		}

		return true
	}

	return false
}

func explode(pair *Pair) (bool) {

	if getDepth(0, pair) < 4 {
		return false
	}

	_,exploded := getFirstDepth4(0,pair)

	literalList := []*int{}
	
	traverseLiterals(pair, func(at *int){
		literalList = append(literalList, at)
	})

	var xChange *int
	var yChange *int

	for at, literal := range literalList {

		if literal == &exploded.xLiteral {
			if (at - 1) >= 0 {
				xChange = literalList[at-1]
			}
		}

		if literal == &exploded.yLiteral {
			if (at + 1) < len(literalList) {
				 yChange = literalList[at+1]
			}
		}
	}

	// set left value
	if xChange != nil {
		*xChange = *xChange + exploded.xLiteral
	}

	// set right value
	if yChange != nil {
		*yChange = *yChange + exploded.yLiteral
	}

	// replace with 0
	parent := exploded.parent

	if parent.xPair == exploded {
		parent.xPair = nil
	}

	if parent.yPair == exploded {
		parent.yPair = nil
	}

	return true
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

func traverse(p *Pair, atFunc func(at *Pair) bool) {
	if !atFunc(p) {
		return
	}

	if p.xPair != nil {
		traverse(p.xPair, atFunc)
	}

	if p.yPair != nil {
		traverse(p.yPair, atFunc)
	}
}

func traverseLiterals(p *Pair, atFunc func(at *int)) {
	if p.xPair == nil {
		atFunc(&p.xLiteral)
	} else {
		traverseLiterals(p.xPair, atFunc)
	}

	if p.yPair == nil {
		atFunc(&p.yLiteral)
	} else {
		traverseLiterals(p.yPair, atFunc)
	}
}
