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

type Report [][]string

func main() {

	// Get Data
	_, file, _, _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	report := Report{}

	for scanner.Scan() {
		numbers := scanner.Text()

		digits := strings.Split(numbers, "")

		report = append(report, digits)
	}

	counts := getCounts(report)

	// Part 1
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

	// Part 2

	// Get oxygen
	var oxygen int64
	oxyReport := report
	reportLen := len(oxyReport)
	for i := 0;i < reportLen;i++  {
		oxyReport = filter(oxyReport, i, "1")

		if len(oxyReport) == 1 {
			oxygen,_ = strconv.ParseInt(strings.Join(oxyReport[0],""), 2, 64)
			break
		}
	}

	var co2 int64
	co2Report := report
	co2ReportLen := len(co2Report)
	for i := 0;i < co2ReportLen;i++  {
		co2Report = filter(co2Report, i, "0")

		if len(co2Report) == 1 {
			co2,_ = strconv.ParseInt(strings.Join(co2Report[0],""), 2, 64)
			break
		}
	}

	fmt.Printf("Part 1: %d\n", gammaInt * epsilonInt)
	fmt.Printf("Part 2: %d\n", oxygen * co2)
}

func filter(report Report, position int, tie string) Report {
	filtered := Report{}
	counts := getCounts(report)
	max := getMax(counts[position], tie)
	min := getMin(counts[position], tie)

	for _,value := range report {

		if (tie == "1") {
			if value[position] == max {
				filtered = append(filtered, value)
			}
		}

		if (tie == "0") {
			if value[position] == min {
				filtered = append(filtered, value)
			}
		}
	}

	return filtered
}

func getMax(counts Counts, tie string) string {

	if counts.Ones == counts.Zeros {
		return tie
	}

	if counts.Ones > counts.Zeros {
		return "1"
	} else {
		return "0"
	}
}

func getMin(counts Counts, tie string) string {

	if counts.Ones == counts.Zeros {
		return tie
	}

	if counts.Ones < counts.Zeros {
		return "1"
	} else {
		return "0"
	}
}

func getCounts(report Report) (counts []Counts) {
	for _,digits := range report {
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

	return counts
}
