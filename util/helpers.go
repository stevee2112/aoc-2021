package util

import (
	"strconv"
	"strings"
	"sort"
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
