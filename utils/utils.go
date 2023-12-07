package utils

import (
	"fmt"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}

	return file
}
