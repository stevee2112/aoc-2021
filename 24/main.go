package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strings"
	"strconv"
)

type Vars struct {
	divider int
	modder int
	adder int
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	pInput, _ := os.Open(path.Dir(file) + "/input")

	defer pInput.Close()
	scanner := bufio.NewScanner(pInput)

	actions := []string{}

	for scanner.Scan() {

		line := scanner.Text()
		actions = append(actions, line)
	}

	vars := []Vars{}

	for i := 0;i < len(actions); i += 18 {
		divParts := strings.Split(actions[i + 4], " ")
		modParts := strings.Split(actions[i + 5], " ")
		addParts := strings.Split(actions[i + 15], " ")
		vars = append(vars, Vars{util.Atoi(divParts[2]), util.Atoi(modParts[2]), util.Atoi(addParts[2])})
	}

	zSet := map[int]string{
		0:"",
	}

	for _,aVar := range vars {
		zSet = newMaxAluActionSet(zSet, aVar.divider, aVar.modder, aVar.adder)
	}

	fmt.Printf("Part 1: %s\n", zSet[0])

	zSet = map[int]string{
		0:"",
	}

	for _,aVar := range vars {
		zSet = newMinAluActionSet(zSet, aVar.divider, aVar.modder, aVar.adder)
	}

	fmt.Printf("Part 2: %s\n", zSet[0])
}

func newMaxAluActionSet(zSet map[int]string, divider int, modAdd int, addAdd int) map[int]string {

	newZ := map[int]string{}

	for z,at := range zSet {
		for w := 9; w >= 1;w-- {
			val := newAluAction(z, w, divider, modAdd, addAdd)
			if currentAt,exists := newZ[val]; exists {
				newAt := at + strconv.Itoa(w)
				if util.Atoi(newAt) > util.Atoi(currentAt) { // Change this fo
					newZ[val] = at + strconv.Itoa(w)
				}
			} else {
				newZ[val] = at + strconv.Itoa(w)
			}
		}
	}

	return newZ
}

func newMinAluActionSet(zSet map[int]string, divider int, modAdd int, addAdd int) map[int]string {

	newZ := map[int]string{}

	for z,at := range zSet {
		for w := 9; w >= 1;w-- {
			val := newAluAction(z, w, divider, modAdd, addAdd)
			if currentAt,exists := newZ[val]; exists {
				newAt := at + strconv.Itoa(w)
				if util.Atoi(newAt) < util.Atoi(currentAt) {
					newZ[val] = at + strconv.Itoa(w)
				}
			} else {
				newZ[val] = at + strconv.Itoa(w)
			}
		}
	}

	return newZ
}

	func newAluAction(z int, w int, divider int, modAdd int, addAdd int) int {

	newVal := 0

	if ((z % 26) + modAdd) == w {
		newVal =  z / divider
	} else {
		newVal = z / divider
		newVal = newVal * 26
		newVal = newVal + (w + addAdd)
	}

	return newVal
}
