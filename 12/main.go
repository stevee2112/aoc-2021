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

	graph := util.Graph{}

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line,"-")

		nodeA := util.GraphNode{}
		nodeB := util.GraphNode{}

		if !graph.NodeExists(nodes[0]) {
			nodeA = util.MakeNode(nodes[0], nil)
			graph.AddNode(nodeA)
		}

		if !graph.NodeExists(nodes[1]) {
			nodeB = util.MakeNode(nodes[1], nil)
			graph.AddNode(nodeB)
		}

		graph.ConnectNodes(nodes[0], nodes[1])
	}

	// Part 1
	part1Count := 0
	graph.Traverse(
		"start",
		func(node util.GraphNode, path []string) bool {

			if path[len(path) - 1] == "end" {
				part1Count++
			}

			return true
		},
		func(node util.GraphNode, path []string) bool {

			if node.Id == "start" {
				return false
			}

			visits := map[string]int{}

			for _,part := range path {
				visits[part]++
			}

			if strings.ToUpper(node.Id) != node.Id { // lowercase
				if visits[node.Id] > 0 {
					return false
				}
			}

			return true // uppercase
		},
	);

	// Part 2
	part2Count := 0
	graph.Traverse(
		"start",
		func(node util.GraphNode, path []string) bool {

			if path[len(path) - 1] == "end" {
				part2Count++
				return false
			}

			return true
		},
		func(node util.GraphNode, path []string) bool {

			if node.Id == "start" {
				return false
			}

			visits := map[string]int{}
			hasDoubleVisit := ""
			tooManyDoubles := false

			for _,part := range path {
				visits[part]++

				if strings.ToUpper(part) != part { // lowercase
					if visits[part] > 1 {
						if visits[part] > 2 {
							return false
						}

						if hasDoubleVisit == "" {
							hasDoubleVisit = part
						} else {
							if hasDoubleVisit != part {
								tooManyDoubles = true
							}
						}
					}
				}

			}

			if strings.ToUpper(node.Id) != node.Id { // lowercase
				if tooManyDoubles {
					return false
				}
			}

			return true // uppercase
		},
	);

	fmt.Printf("Part 1: %d\n", part1Count)
	fmt.Printf("Part 2: %d\n", part2Count)
}
