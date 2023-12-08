package utils

import (
	"fmt"
	"os"
	"strconv"
)

func OpenFile(path string) *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}

	return file
}

func StringsToInts(line []string) []int {
	var ints []int
	for _, element := range line {
		ints = append(ints, StringToInt(element))
	}
	return ints
}

func StringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return value
}
