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

type Counts struct {
	Ones int
	Zeros int
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	counts := []Counts{}

	for scanner.Scan() {
		numbers := scanner.Text()
		digits := strings.Split(numbers, "")

		for i, val := range digits {
			if (i + 1) > len(counts) {
				counts = append(counts, Counts{0,0})
			}

			if val == "1" {
				counts[i].Ones++
			}

			if val == "0" {
				counts[i].Zeros++
			}
		}
	}

	gamma := "";
	epsilon := "";

	for _,pos := range counts {
		if pos.Ones > pos.Zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt,_ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt,_ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("Part 1: %d\n", gammaInt * epsilonInt)
	fmt.Printf("Part 2: %d\n", 0)
}
