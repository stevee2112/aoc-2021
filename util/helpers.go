package util

import (
	"strconv"
	"strings"
	"sort"
	"fmt"
)

// Atoi as oneliner as we fully control the input
func Atoi(string string) int {
	intVal, _ := strconv.Atoi(string)
	return intVal
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func HexToBin(hex string) (string) {
	binString := ""

	for _,hexChar := range strings.Split(hex, "") {
		ui, _ := strconv.ParseUint(hexChar, 16, 64)
		binString += strings.Join(strings.Split(fmt.Sprintf("%016b", ui), "")[12:16], "")
	}

	return binString
}

func CloneIntMap(cloner map[int]int) map[int]int {
	clone := map[int]int{}
	for k,v := range cloner {
	  clone[k] = v
	}

	return clone
}

func CloneStringIntMap(cloner map[string]int) map[string]int {
	clone := map[string]int{}
	for k,v := range cloner {
	  clone[k] = v
	}

	return clone
}

func UniqueIntSlice(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}	
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}
