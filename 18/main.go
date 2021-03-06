package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
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

	pairs := []*Pair{}
	pairStrings := []string{}

	for scanner.Scan() {
		pairStr := scanner.Text()
		pairs = append(pairs, parsePair(pairStr))
		pairStrings = append(pairStrings, pairStr)
	}

	var pair *Pair
	for _, current := range pairs {

		if pair == nil {
			pair = current
			continue
		}

		pair = addPairs(pair, current)
		reduce(pair)
	}

	part1 := computeMagnitude(pair)

	max := 0

	for i := 0; i < len(pairStrings);i++ {
		for j := 0; j < len(pairStrings);j++ {
			if i == j {
				continue
			}

			sum1 := addPairs(parsePair(pairStrings[i]),parsePair(pairStrings[j]))
			sum2 := addPairs(parsePair(pairStrings[j]),parsePair(pairStrings[i]))

			reduce(sum1)
			reduce(sum2)

			ab := computeMagnitude(sum1)
			ba := computeMagnitude(sum2)

			if ab > max {
				max = ab
			}

			if ba > max {
				max = ba
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", max)
}

func computeMagnitude(pair *Pair) int {
	sum := 0
	if pair.xPair == nil {
		sum += 3 * pair.xLiteral
	} else {
		sum += 3 * computeMagnitude(pair.xPair)
	}

	if pair.yPair == nil {
		sum += 2 * pair.yLiteral
	} else {
		sum += 2 * computeMagnitude(pair.yPair)
	}

	return sum
}

func reduce(pair *Pair) {
	for {
		exploded :=explode(pair)

		if exploded {
			continue
		}

		splited :=split(pair)

		if splited {
			continue
		}

		break
	}

	return
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

	literalList := []struct{
		literal *int
		pair *Pair
	}{}

	traverseLiterals(pair, func(at *int, pair *Pair){
		literalList = append(literalList, struct{
			literal *int
			pair *Pair
		} {
			at,
			pair,
		})
	})

	for _,pair := range literalList {

		if *pair.literal >= 10 {
			if pair.literal == &pair.pair.xLiteral {
				isLeft = true
			}

			if pair.literal == &pair.pair.yLiteral {
				isLeft = false
			}

			toSplit = pair.pair
			break
		}
	}

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

	traverseLiterals(pair, func(at *int, pair *Pair){
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
			if (at + 1) < len(literalList)  {
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
	a.parent = pair
	b.parent = pair
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

func traverseLiterals(p *Pair, atFunc func(at *int, pair *Pair)) {
	if p.xPair == nil {
		atFunc(&p.xLiteral, p)
	} else {
		traverseLiterals(p.xPair, atFunc)
	}

	if p.yPair == nil {
		atFunc(&p.yLiteral, p)
	} else {
		traverseLiterals(p.yPair, atFunc)
	}
}
