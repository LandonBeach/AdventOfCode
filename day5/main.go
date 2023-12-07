package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
)

// Thoughts:
// Create Seed struct containing soil, fertilizer, water, light, temperature, humidity, and location values
// Create the "maps" for each stage transition. Maybe a struct with methods?
//   Destination, Source, Range
//   Delta (dest minus source) to convert from one value to the other value

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
