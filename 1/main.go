package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	lastDepth := -1
	increaseCounter := 0
	depthSums := []int{}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	at := 0
	for scanner.Scan() {

		at++
		currentDepth,_ := strconv.Atoi(scanner.Text())

		if (at - 1) >= 0 {
			if ((at - 1) + 1) > len(depthSums) {
				depthSums = append(depthSums, currentDepth)
			}
		}

		if (at - 2) >= 0 {
			depthSums[at - 2] = depthSums[at - 2] + currentDepth
		}

		if (at - 3) >= 0 {
			depthSums[at - 3] = depthSums[at - 3] + currentDepth
		}

		if lastDepth == -1 {
			lastDepth = currentDepth
			continue
		}

		if currentDepth > lastDepth {
			increaseCounter++
		}

		lastDepth = currentDepth

	}

	lastDepthSum := -1
	increaseCounterDepthSums := 0

	for _,sum := range depthSums {

		if lastDepthSum == -1 {
			lastDepthSum = sum
			continue
		}

		if sum > lastDepthSum {
			increaseCounterDepthSums++
		}

		lastDepthSum = sum
	}

	fmt.Printf("Part 1: %d\n", increaseCounter)
	fmt.Printf("Part 2: %d\n", increaseCounterDepthSums)
}
