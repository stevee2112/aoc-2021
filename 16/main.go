package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"stevee2112/aoc-2021/util"
	"strconv"
)

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	//data := []int{}

	packet := ""
	for scanner.Scan() {
		line := scanner.Text()
		packet = util.HexToBin(line)
	}

	part1VersionSum, _, result  :=  parsePacket(packet)

	fmt.Printf("Part 1: %d\n", part1VersionSum)
	fmt.Printf("Part 2: %d\n", result)
}

func parsePacket(packet string) (int, string, int) {

	sequence := strings.Split(packet, "")
	startPacket := true
	currentPacketVersion := 0
	currentPacketType := 0
	bitsInPacket := 0
	literal := ""

	for len(sequence) > 0 {
		if startPacket == true {
			packetVersion,_ := strconv.ParseInt(strings.Join(sequence[0:3],""), 2, 64)
			packetType,_ := strconv.ParseInt(strings.Join(sequence[3:6],""), 2, 64)

			currentPacketVersion = int(packetVersion)
			currentPacketType = int(packetType)

			sequence = sequence[6:]

			startPacket = false

			// Reset stuff
			bitsInPacket = 6 // read so far
			literal = ""
		} else {
			switch currentPacketType {
			case 4: // literal
				// Get group
				group := sequence[0:5]
				bitsInPacket += len(group)

				sequence = sequence[5:]

				lastGroup := false
				literal += strings.Join(group[1:5],"")

				if group[0] == "0" {
					lastGroup = true
				}

				if lastGroup {
					literalInt,_ := strconv.ParseInt(literal, 2, 64)
					return currentPacketVersion, strings.Join(sequence,""), int(literalInt)
				}
			default: // operator
				id := sequence[0]
				sequence = sequence[1:]

				lengthOrCountBits := 0

				if id == "0" { // length 15
					lengthOrCountBits = 15
				}

				if id == "1" { // length 11
					lengthOrCountBits = 11
				}

				lengthOrCount,_  := strconv.ParseInt(strings.Join(sequence[0:lengthOrCountBits], ""), 2, 64)

				sequence = sequence[lengthOrCountBits:]


				versionSum := 0
				operand := 0
				operands := []int{}
				if id == "0" {

					subversion := 0
					subpackets := strings.Join(sequence[0:lengthOrCount],"")
					sequence = sequence[lengthOrCount:]
					for len(subpackets) > 0 {
						subversion, subpackets, operand = parsePacket(subpackets)
						operands = append(operands, operand)
						versionSum += subversion
					}

					versionSum += currentPacketVersion
				}

				if id == "1" { // pass whole remaining sequence lengthOrCount amount of times
					subversion := 0
					subsequence := strings.Join(sequence,"")
					for i := 0; i < int(lengthOrCount);i++ {
						subversion, subsequence, operand = parsePacket(subsequence)
						operands = append(operands, operand)
						versionSum += subversion
					}

					sequence = strings.Split(subsequence,"")
					versionSum += currentPacketVersion
				}

				result := 0
				switch currentPacketType {
				case 0: // SUM
					for _,val := range operands {
						result += val
					}
				case 1: // SUM
					result = 1
					for _,val := range operands {
						result *= val
					}
				case 2: // MIN
					result = operands[0]
					for _,val := range operands {
						if val < result {
							result = val
						}
					}
				case 3: // MAX
					result = operands[0]
					for _,val := range operands {
						if val > result {
							result = val
						}
					}
				case 5: // GT
					result = 0
					if operands[0] > operands[1] {
						result = 1
					}
				case 6: // LT
					result = 0
					if operands[0] < operands[1] {
						result = 1
					}
				case 7: // EQUAL
					result = 0
					if operands[0] == operands[1] {
						result = 1
					}
				}


				return versionSum, strings.Join(sequence, ""), result
			}
		}
	}

	// Shoul never really get here
	return 0, "", 1
}
