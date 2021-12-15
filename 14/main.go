package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type Polymer struct {
	Source []string
	Rules map[string]string
	Counts map[string]int
}

func NewPolymer(source string, rules map[string]string) Polymer {
	polymer := Polymer{}
	polymer.Source = strings.Split(source, "")
	polymer.Rules = rules
	polymer.Counts = map[string]int{}

	return polymer
}

func (p Polymer) String() string {
	return strings.Join(p.Source,"")
}

func (p *Polymer) Step() {

	p.Counts = map[string]int{}
	at := 0
	done := false
	for !done {

		tuple := p.Source[at] + p.Source[at + 1]

		p.Counts[p.Source[at]]++

		if ele,exists := p.Rules[tuple]; exists {

			insert := at + 1
			p.Source = append(p.Source, "")
			copy(p.Source[(insert+1):], p.Source[insert:])
			p.Source[insert] = ele

			p.Counts[ele]++

			at = at + 2
		} else {
			at++
		}

		if at >= len(p.Source) - 1 {

			// add last char
			p.Counts[p.Source[len(p.Source) - 1]]++
			done = true
		}
	}
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
	for i := 0;i < steps;i++ {
		polymer.Step()
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
	fmt.Printf("Part 2: %d\n", 0)
}
