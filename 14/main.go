package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type PolymerTuple struct {
	Tuple string
	Created int
}

type Polymer struct {
	Tuples map[string]int
	Rules map[string]string
	Counts map[string]int
}

func NewPolymer(source string, rules map[string]string) Polymer {
	polymer := Polymer{}
	polymer.Tuples = map[string]int{}
	polymer.Rules = rules
	polymer.Counts = map[string]int{}

	lastChar := ""
	for _,char := range strings.Split(source, "") {
		polymer.Counts[char]++
		if lastChar == "" {
			lastChar = char
			continue
		}

		tuple := PolymerTuple{
			Tuple: lastChar + char,
			Created: 0,
		}

		lastChar = char

		polymer.Tuples[tuple.Tuple]++
	}

	return polymer
}

func (p *Polymer) Step(id int) {

	newTuples := map[string]int{}
	for k,v := range p.Tuples {
		newTuples[k] = v
	}

	for source,result := range p.Rules {

		if count,exists := p.Tuples[source]; exists {

			aKey := source[0:1] + result
			bKey := result + source[1:2]

			newTuples[aKey] += count
			newTuples[bKey] += count
			newTuples[source] -= count
			p.Counts[result] += count
		}
	}

	p.Tuples = newTuples

}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	source := ""
	rulesSection := false
	rules := map[string]string{}

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			rulesSection = true
			continue
		}

		if !rulesSection {
			source = line
		}

		if rulesSection {
			parts := strings.Split(line, " -> ")
			rules[parts[0]] = parts[1]
		}
	}

	polymer := NewPolymer(
		source,
		rules,
	)

    // Part 1
    steps := 10
    for i := 1;i <= steps;i++ {
		polymer.Step(i)
    }

	counts := polymer.Counts
	max := 0
	min := 0

	for _,val := range counts {
		if max == 0 || val > max {
			max = val
		}

		if min == 0 || val < min {
			min = val
		}
	}

	fmt.Printf("Part 1: %d\n", max - min)

	// Part 2
    steps = 30 // Do 30 more steps
    for i := 1;i <= steps;i++ {
		polymer.Step(i)
    }

	counts = polymer.Counts
	max = 0
	min = 0

	for _,val := range counts {
		if max == 0 || val > max {
			max = val
		}

		if min == 0 || val < min {
			min = val
		}
	}


	fmt.Printf("Part 2: %d\n", max - min)
}
