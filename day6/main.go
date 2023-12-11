package main

import (
	"AdventOfCode/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Time: <int> <int>
	// Distance: <int> <int>
	file := strings.Split(string(fileImport), "\n")
	fileTimes := strings.Split(file[0], ":")[1]
	fileDistance := strings.Split(file[1], ":")[1]

	// Convert the string to []int
	raceTime := utils.StringsToInts(strings.Fields(fileTimes))
	raceDistance := utils.StringsToInts(strings.Fields(fileDistance))

	var recordBreakingWays []int

	for race, recordTime := range raceTime {
		waysToBeat := 0
		for millisecond := 1; millisecond < recordTime; millisecond++ {
			remainingTime := recordTime - millisecond
			if distance := remainingTime * millisecond; distance > raceDistance[race] {
				waysToBeat++
			}
		}
		recordBreakingWays = append(recordBreakingWays, waysToBeat)
	}

	part1Total := 1
	for _, numberWays := range recordBreakingWays {
		part1Total *= numberWays
	}

	fmt.Println("Part 1 - Number of ways you can beat the record all multiplied together:", part1Total)
}
