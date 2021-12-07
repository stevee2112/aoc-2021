package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
	"sort"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	in := strings.Split(scanner.Text(),",")

	numbers := []int{}

	for _,numberStr := range in {
		numbers = append(numbers, util.Atoi(numberStr))
	}

	sort.Ints(numbers)

	costPart1 := 9999999999999

	for _,at := range numbers {
		cost := computeCostPart1(at, numbers)

		if cost > costPart1 {
			break
		} else {
			costPart1 = cost
		}
	}

	fmt.Printf("Part 1: %d\n", costPart1)
	fmt.Printf("Part 2: %d\n", 0)
}

func computeCostPart1(position int, numbers []int) int {

	cost := 0

	for _,at := range numbers {
		cost += util.Abs(position - at)
	}

	return cost
}
