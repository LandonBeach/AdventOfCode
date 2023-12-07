package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	runningSum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		runningSum += firstLastDigits(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	fmt.Println(runningSum)
}

// Gets the first and last digit in the string and returns a single int
func firstLastDigits(line string) int {
	var firstDigit, lastDigit int
	var firstIndex, lastIndex int = len(line), -1

	// Iterate forwards to get first digit
	for i := 0; i < len(line); i++ {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			firstDigit = digit
			firstIndex = i
			break
		}
	}

	// Iterate backwards to get second digit
	for i := len(line) - 1; i >= 0; i-- {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			lastDigit = digit
			lastIndex = i
			break
		}
	}

	// Digit word to value map
	numberSubStrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	// Iterate through the digit words and find occurrences within the string/line
	for subStr, value := range numberSubStrings {
		// Search for the first digit word.
		if index := strings.Index(line, subStr); index >= 0 {
			if index < firstIndex {
				firstIndex = index
				firstDigit = value
			}
		}

		// Search for the last digit word.
		if index := strings.LastIndex(line, subStr); index >= 0 {
			if index > lastIndex {
				lastIndex = index
				lastDigit = value
			}
		}
	}

	// Combine both digits into a single int
	return (firstDigit * 10) + lastDigit
}
