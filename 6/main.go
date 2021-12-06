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
	initialPopulation := strings.Split(scanner.Text(),",")
	model := map[int]int{
		0:0,
		1:0,
		2:0,
		3:0,
		4:0,
		5:0,
		6:0,
		7:0,
		8:0,
	}

	for _,fish := range initialPopulation {
		intFish := util.Atoi(fish)
		model[intFish] = model[intFish] + 1
	}

	days := 256

	part1Count := 0
	part2Count := 1
	for i := 1;i <= days;i++ {
		model = spawn(model)

		if i == 80 {
			part1Count = count(model)
		}

		if i == 256 {
			part2Count = count(model)
		}
	}

	fmt.Printf("Part 1: %d\n", part1Count)
	fmt.Printf("Part 2: %d\n", part2Count)

}

func spawn(model map[int]int) map[int]int {

	newModel := map[int]int{
		0:0,
		1:0,
		2:0,
		3:0,
		4:0,
		5:0,
		6:0,
		7:0,
		8:0,
	}

	newModel[0] = model[1]
	newModel[1] = model[2]
	newModel[2] = model[3]
	newModel[3] = model[4]
	newModel[4] = model[5]
	newModel[5] = model[6]
	newModel[6] = model[7]
	newModel[7] = model[8]
	newModel[6] += model[0] // Parents start over
	newModel[8] = model[0] // Parents make baby

	return newModel
}

func count(model map[int]int) int {

	sum := 0
	for _,age := range model {
		sum += age
	}

	return sum
}
