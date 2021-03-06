package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
	"unicode"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	burrow := util.Grid{}

	y := 0

	seenMap := map[string]int{}
	atMap := map[string]*util.Coordinate{}

	for scanner.Scan() {

		line := scanner.Text()

		x := 0
		for _,val := range strings.Split(line,"") {

			displayVal := val
			isLetter := false
			atNum := 0
			if unicode.IsLetter([]rune(val)[0]) {

				isLetter = true
				if i, seen := seenMap[val]; seen {
					atNum = i + 1
				}

				displayVal = fmt.Sprintf("%s%d", strings.ToLower(val), atNum)

			}
			burrow.SetValue(x, y, displayVal)

			if isLetter {
				at := burrow.GetCoordinate(x, y)
				atMap[displayVal] = &at
				seenMap[val] = atNum
			}

			x++

		}
		y++
	}

	energy := 0
	lastMove := []string{}

	for {
		burrow.PrintGrid(3)

		if done(burrow) {
			break
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Energy Used (%d)\n\n", energy)
		fmt.Print("Move: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		undo := false
		if text == "u" {

			if len(lastMove) < 1 {
				continue
			}

			undo = true
			text = lastMove[0]
			lastMove = lastMove[1:]
		}
		parts := strings.Split(text, " ")

		if len(parts) < 3 {
			continue
		}

		cost := 0

		burrow, atMap, cost = move(burrow, parts[0], parts[1], util.Atoi(parts[2]), atMap, undo)
		energy += cost

		if !undo {
			lastMove = append([]string{text}, lastMove...)
		}
	}

	for _, move := range lastMove {
		fmt.Println(move)
	}

	fmt.Printf("Total Energy: %d\n", energy)
}

func move(
	burrow util.Grid,
	toMove string,
	direction string,
	distance int,
	atMap map[string]*util.Coordinate,
	undo bool,
) (util.Grid, map[string]*util.Coordinate, int) {

	at, found := atMap[toMove]; if !found {
		return burrow, atMap, 0
	}

	if undo {
		switch (direction) {
		case "u":
			direction = "d"
		case "d":
			direction = "u"
		case "l":
			direction = "r"
		case "r":
			direction = "l"
		}
	}

	// Clear current value
	burrow.SetValue(at.X, at.Y, ".")

	// Set new value
	switch (direction) {
	case "u":
		burrow.SetValue(at.X, at.Y - distance, toMove)
		coor := burrow.GetCoordinate(at.X, at.Y - distance)
		atMap[toMove] = &coor
	case "d":
		burrow.SetValue(at.X, at.Y + distance, toMove)
		coor := burrow.GetCoordinate(at.X, at.Y + distance)
		atMap[toMove] = &coor
	case "l":
		burrow.SetValue(at.X - distance, at.Y, toMove)
		coor := burrow.GetCoordinate(at.X - distance, at.Y)
		atMap[toMove] = &coor
	case "r":
		burrow.SetValue(at.X + distance, at.Y, toMove)
		coor := burrow.GetCoordinate(at.X + distance, at.Y)
		atMap[toMove] = &coor
	}

	cost := 0

	char := strings.ToUpper(string(toMove[0]))
	switch (char) {
	case "A":
		cost += (distance * 1)
	case "B":
		cost += (distance * 10)
	case "C":
		cost += (distance * 100)
	case "D":
		cost += (distance * 1000)
	}

	if undo {
		cost *= -1
	}

	return burrow, atMap, cost
}

func done(burrow util.Grid) bool {

	done := true
	if string(strings.ToUpper(burrow.GetCoordinate(3,2).Value.(string))[0]) != "A" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(3,3).Value.(string))[0]) != "A" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(3,4).Value.(string))[0]) != "A" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(3,5).Value.(string))[0]) != "A" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(5,2).Value.(string))[0]) != "B" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(5,3).Value.(string))[0]) != "B" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(5,4).Value.(string))[0]) != "B" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(5,5).Value.(string))[0]) != "B" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,2).Value.(string))[0]) != "C" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,3).Value.(string))[0]) != "C" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,4).Value.(string))[0]) != "C" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,5).Value.(string))[0]) != "C" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(9,2).Value.(string))[0]) != "D" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(9,3).Value.(string))[0]) != "D" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,4).Value.(string))[0]) != "C" {
		return false
	}

	if string(strings.ToUpper(burrow.GetCoordinate(7,5).Value.(string))[0]) != "C" {
		return false
	}

	return done
}
