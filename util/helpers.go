package util

import (
	"strconv"
)

// Atoi as oneliner as we fully control the input
func Atoi(string string) int {
	intVal, _ := strconv.Atoi(string)
	return intVal
}
