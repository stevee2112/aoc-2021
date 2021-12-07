package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
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

	min := 999999999999999999
	max := 0

	for _,numberStr := range in {
		number :=  util.Atoi(numberStr)

		if number < min {
			min = number
		}

		if number > max {
			max = number
		}
		numbers = append(numbers, util.Atoi(numberStr))
	}

	costPart1 := 9999999999999
	costPart2 := 9999999999999

	for at := min;at <= max;at++ {
		cost := computeCostPart1(at, numbers)

		if cost > costPart1 {
			break
		} else {
			costPart1 = cost
		}
	}

	for at := min;at <= max;at++ {
		cost := computeCostPart2(at, numbers)

		if cost > costPart2 {
			break
		} else {
			costPart2 = cost
		}
	}

	fmt.Printf("Part 1: %d\n", costPart1)
	fmt.Printf("Part 2: %d\n", costPart2)
}

func computeCostPart1(position int, numbers []int) int {

	cost := 0

	for _,at := range numbers {
		cost += util.Abs(position - at)
	}

	return cost
}

func computeCostPart2(position int, numbers []int) int {

	cost := 0

	for _,at := range numbers {
		diff := util.Abs(position - at)
		cost += ((diff * diff) + diff) / 2
	}

	return cost
}

