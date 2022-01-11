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
	"unicode"
	//"math"
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

	fmt.Println(vars)

	// at := 99999999999

	// skipped :=0
	// i := at
	// for i > 0 {
	// 	registry := map[string]int{
	// 		"w": 0,
	// 		"x": 0,
	// 		"y": 0,
	// 		"z": 0,
	// 	}

	// 	asString := strconv.Itoa(i)

	// 	if strings.Contains(asString, "0") {
	// 		skip := int(math.Pow(float64(10), float64((len(asString) - strings.Index(asString, "0")) - 1)))
	// 		i -= skip
	// 		continue
	// 	}

	// 	input := strings.Split(asString, "")

	// 	//fmt.Println("NEW", asString)
	// 	for _, line := range actions {
	// 		parts := strings.Split(line, " ")
	// 		//fmt.Println(line)

	// 		if parts[0] == "inp" {
	// 			fmt.Println(registry, input)
	// 		}

	// 		registry, input = doAluAction(registry, input, parts[0], parts[1:])
	// 	}

	// 	//fmt.Println(i, registry, input)
	// 	if registry["z"] == 0 {
	// 		fmt.Println("MATCH", i, registry)
	// 		break
	// 	}

	// 	skipped++
	// 	i--
	// }

	zSet := map[int]string{
		0:"",
	}

	fmt.Println("newWay")

	//fmt.Println(newAluAction(12, 9, 1, 14, 3))

	for i,aVar := range vars {
		fmt.Println(i, len(zSet))
		zSet = newAluActionSet(zSet, aVar.divider, aVar.modder, aVar.adder)
		//fmt.Println(zSet)
	}

	// zSet = newAluAction(zSet, 26, -5, 13)
	// fmt.Println(zSet)
	// zSet = newAluAction(zSet, 26, -8, 3)
	// fmt.Println(zSet)
	// zSet = newAluAction(zSet, 26, -11, 10)

	fmt.Println(len(zSet))
	fmt.Println(zSet[0])

}

func newAluActionSet(zSet map[int]string, divider int, modAdd int, addAdd int) map[int]string {

	newZ := map[int]string{}
    // if (z mod 26) + modAdd == w)
	//     z = z / 26, 
	// else 
	// 	z = z + w + 5

	for z,at := range zSet {
		for w := 9; w >= 1;w-- {
			val := newAluAction(z, w, divider, modAdd, addAdd)
			if currentAt,exists := newZ[val]; exists {
				newAt := at + strconv.Itoa(w)
				if util.Atoi(newAt) > util.Atoi(currentAt) {
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


	//fmt.Println("new", z, w, divider, modAdd, addAdd)
    // if (z mod 26) + modAdd == w)
	//     z = z / 26, 
	// else 
	// 	z = z + w + 5

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


func doAluAction(registry map[string]int, input []string, instruction string, variables []string) (map[string]int, []string) {

	switch (instruction) {
	case "inp":
		registry[variables[0]] = util.Atoi(input[0])
		input = input[1:]
	case "add":
		var val int
		if unicode.IsLetter([]rune(variables[1])[0]) {
			val = registry[variables[1]]
		} else {
			val = util.Atoi(variables[1])
		}
		registry[variables[0]] = registry[variables[0]] + val
	case "mul":
		var val int
		if unicode.IsLetter([]rune(variables[1])[0]) {
			val = registry[variables[1]]
		} else {
			val = util.Atoi(variables[1])
		}
		registry[variables[0]] = registry[variables[0]] * val
	case "div":
		var val int
		if unicode.IsLetter([]rune(variables[1])[0]) {
			val = registry[variables[1]]
		} else {
			val = util.Atoi(variables[1])
		}
		registry[variables[0]] = registry[variables[0]] / val
	case "mod":
		var val int
		if unicode.IsLetter([]rune(variables[1])[0]) {
			val = registry[variables[1]]
		} else {
			val = util.Atoi(variables[1])
		}
		registry[variables[0]] = registry[variables[0]] % val
	case "eql":
		var val int
		if unicode.IsLetter([]rune(variables[1])[0]) {
			val = registry[variables[1]]
		} else {
			val = util.Atoi(variables[1])
		}

		if registry[variables[0]] == val {
			registry[variables[0]] = 1
		} else {
			registry[variables[0]] = 0
		}
	}

	return registry, input
}
