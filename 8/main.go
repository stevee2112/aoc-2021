package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"stevee2112/aoc-2021/util"
	"strconv"
	"strings"
)

type NoteEntry struct {
	Pattern []string
	Output []string
}

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	notes := []NoteEntry{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")

		noteEntry := NoteEntry{
			Pattern: strings.Split(parts[0], " "),
			Output: strings.Split(parts[1], " "),
		}

		notes = append(notes, noteEntry)
	}

	// Part 1
	uniqueLenDigits := 0
	for _,note := range notes {

		for _,output := range note.Output {
			segmentCount := len(output)

			if segmentCount == 2 { // 1
				uniqueLenDigits++
			}

			if segmentCount == 4 { // 4
				uniqueLenDigits++
			}

			if segmentCount == 3 { // 7
				uniqueLenDigits++
			}

			if segmentCount == 7 { // 8
				uniqueLenDigits++
			}
		}
	}

	// Part 2
	sum := 0
	for _,note := range notes {
		sum += computeOutput(note)
	}

	fmt.Printf("Part 1: %d\n", uniqueLenDigits)
	fmt.Printf("Part 2: %d\n", sum)
}

func computeOutput(note NoteEntry) int {
	lettersToPosition := map[string]string{}
	positionToLetter := map[string]string{}

	sequenceToNumber := map[string]int{}
	numberToSequence := map[int]string{}

	// Find 1, 4, 3, 7
	for _,sequence := range note.Pattern {
		segmentCount := len(sequence)

		if segmentCount == 2 { // 1
			sequenceToNumber[util.SortString(sequence)] = 1
			numberToSequence[1] = util.SortString(sequence)
		}

		if segmentCount == 4 { // 4
			sequenceToNumber[util.SortString(sequence)] = 4
			numberToSequence[4] = util.SortString(sequence)
		}

		if segmentCount == 3 { // 7
			sequenceToNumber[util.SortString(sequence)] = 7
			numberToSequence[7] = util.SortString(sequence)

		}

		if segmentCount == 7 { // 8
			sequenceToNumber[util.SortString(sequence)] = 8
			numberToSequence[8] = util.SortString(sequence)
		}
	}

	// Letter in 7 that is NOT 1 is the top segment
	for _,letter := range strings.Split(numberToSequence[7], "") {
		if !strings.Contains(numberToSequence[1], letter) {
			lettersToPosition[letter] = "T"
			positionToLetter["T"] = letter
			break;
		}
	}

	// for each segment with 5 (2,3,5) the segment that has both letters from 1 is 3
	for _,sequence := range note.Pattern {
		segmentCount := len(sequence)

		if segmentCount == 5 {
			lettersInOne := strings.Split(numberToSequence[1],"")

			if strings.Contains(sequence, lettersInOne[0]) && strings.Contains(sequence, lettersInOne[1]) {
				sequenceToNumber[util.SortString(sequence)] = 3
				numberToSequence[3] = util.SortString(sequence)
				break
			}
		}
	}

	// what ever is in 4 AND 3 that is NOT in 7 is the middle (now know middle)
	for _,letter := range strings.Split(numberToSequence[4], "") {
		if strings.Contains(numberToSequence[3], letter) && !strings.Contains(numberToSequence[7], letter) {
			lettersToPosition[letter] = "M"
			positionToLetter["M"] = letter
			break
		}
	}

	// the sequence of 6 that does NOT have the middle sequence is 0
	for _,sequence := range note.Pattern {
		segmentCount := len(sequence)

		if segmentCount == 6 {
			if !strings.Contains(sequence, positionToLetter["M"]) {
				sequenceToNumber[util.SortString(sequence)] = 0
				numberToSequence[0] = util.SortString(sequence)
			}
		}
	}

	// the letter in 4 that is NOT in 1 and is NOT the middle is the top left (now know top left)
	for _,letter := range strings.Split(numberToSequence[4], "") {
		if !strings.Contains(numberToSequence[1], letter) && (letter != positionToLetter["M"]) {
			lettersToPosition[letter] = "TL"
			positionToLetter["TL"] = letter
			break
		}
	}

	// the sequence with 5 segments and has the top left letter is 5
	for _,sequence := range note.Pattern {
		segmentCount := len(sequence)

		if segmentCount == 5 {
			if strings.Contains(sequence, positionToLetter["TL"]) {
				sequenceToNumber[util.SortString(sequence)] = 5
				numberToSequence[5] = util.SortString(sequence)
				break
			}
		}
	}


	// the remaining sequence of 5 is 2
	for _,sequence := range note.Pattern {
		segmentCount := len(sequence)

		if segmentCount == 5 {
			if _, found := sequenceToNumber[util.SortString(sequence)]; !found {
				sequenceToNumber[util.SortString(sequence)] = 2
				numberToSequence[2] = util.SortString(sequence)
				break
			}
		}
	}

	// the letter in 5 that is NOT top and not in 4 is the bottom (now know bottom)
	for _,letter := range strings.Split(numberToSequence[5], "") {
		if letter != positionToLetter["T"] && !strings.Contains(numberToSequence[4], letter) {
			lettersToPosition[letter] = "B"
			positionToLetter["B"] = letter
			break
		}
	}

	// the remain letter in 5 is bottom left (now know bottom right)
	for _,letter := range strings.Split(numberToSequence[5], "") {
		if _, found := lettersToPosition[letter]; !found {
			lettersToPosition[letter] = "BR"
			positionToLetter["BR"] = letter
			break
		}
	}

	// the remaining unknown letter in 3 is top right (now know top right)
	for _,letter := range strings.Split(numberToSequence[3], "") {
		if _, found := lettersToPosition[letter]; !found {
			lettersToPosition[letter] = "TR"
			positionToLetter["TR"] = letter
			break
		}
	}

	// the remaining unknown sequence that has all letters known 9
	for _,sequence := range note.Pattern {
		if _, found := sequenceToNumber[util.SortString(sequence)]; !found {
			hasAll := true
			for _,letter := range strings.Split(sequence, "") {
				if _, letterFound := lettersToPosition[letter]; !letterFound {
					hasAll = false
					break
				}
			}

			if hasAll {
				sequenceToNumber[util.SortString(sequence)] = 9
				numberToSequence[9] = (util.SortString(sequence))
				break
			}
		}
	}

	// the remain sequence is 6
	for _,sequence := range note.Pattern {
		if _, found := sequenceToNumber[util.SortString(sequence)]; !found {
			sequenceToNumber[util.SortString(sequence)] = 6
			numberToSequence[6] = util.SortString(sequence)
			break
		}
	}

	// the last unknown number in 6 is bottom left (now know bottom left)
    for _,letter := range strings.Split(numberToSequence[6], "") {
        if _, found := lettersToPosition[letter]; !found {
            lettersToPosition[letter] = "BL"
            positionToLetter["BL"] = letter
            break
        }
    }

	outputStr := ""

	for _,sequence := range note.Output {
		outputStr += strconv.Itoa(sequenceToNumber[util.SortString(sequence)])
	}

	return util.Atoi(outputStr)
}
