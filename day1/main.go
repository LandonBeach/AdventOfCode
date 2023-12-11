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

	part1Sum := 0
	part2Sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		part1SubTotal, part2SubTotal := firstLastDigits(scanner.Text())
		part1Sum += part1SubTotal
		part2Sum += part2SubTotal
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	fmt.Println("Sum for Part1 solution:", part1Sum)
	fmt.Println("Sum for Part2 solution:", part2Sum)
}

// Gets the first and last digit in the string and returns a single int
func firstLastDigits(line string) (int, int) {
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

	// Combine both digits into a single int
	part1SubTotal := (firstDigit * 10) + lastDigit

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
	part2SubTotal := (firstDigit * 10) + lastDigit

	// Return the sub-totals for the part 1 and part 2 solutions
	return part1SubTotal, part2SubTotal
}
