package main

import (
	"AdventOfCode/utils"
	"fmt"
	"os"
	"strings"
)

/*
 * Day 6 Instructions:
 * You will get a fixed amount of time during which your boat has to travel as far as it can, and you win if your boat goes the farthest.
 * Holding down the button charges the boat, and releasing the button allows the boat to move.
 * Boats move faster if their button was held longer, but time spent holding the button counts against the total race time.
 * You can only hold the button at the start of the race, and boats don't move until the button is released.
 * The boat has a starting speed of zero millimeters per millisecond.
 * For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.
 * To see how much margin of error you have, determine the number of ways you can beat the record in each race and then multiply the results into a single value.
 */
func main() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Part 1 - each column indicates the total time of the race and the record distance
	// Time: <int> <int>
	// Distance: <int> <int>
	file := strings.Split(string(fileImport), "\n")
	fileTimes := strings.Split(file[0], ":")[1]
	fileDistance := strings.Split(file[1], ":")[1]

	// Convert the string to []int
	raceTime := utils.StringsToInts(strings.Fields(fileTimes))
	recordDistance := utils.StringsToInts(strings.Fields(fileDistance))

	recordBreakingWays := findRecordBreakingWays(raceTime, recordDistance)

	part1Total := 1
	for _, numberWays := range recordBreakingWays {
		part1Total *= numberWays
	}

	fmt.Println("Part 1 - Number of ways you can beat the record all multiplied together:", part1Total)

	// Part 2 - Input has bad kerning. There is only 1 race so we need to remove all spaces and make use a single number.
	totalTime := utils.StringToInt(strings.ReplaceAll(fileTimes, " ", ""))
	totalDistance := utils.StringToInt(strings.ReplaceAll(fileDistance, " ", ""))
	recordBreakingWays = findRecordBreakingWays([]int{totalTime}, []int{totalDistance})

	fmt.Println("Part 2 - Number of ways you can beat the record:", recordBreakingWays[0])
}

func findRecordBreakingWays(raceTime []int, recordDistance []int) []int {
	var recordBreakingWays []int

	for race, time := range raceTime {
		waysToBeat := 0
		prevDistance := 0
		for millisecond := recordDistance[race] / time; millisecond < time; millisecond++ {
			remainingTime := time - millisecond
			if distanceTravelled := remainingTime * millisecond; distanceTravelled > recordDistance[race] {

				// Only need to check half since the results are a palindrome because we have constant acceleration and deceleration. For example:
				//   Hold the button for 2 milliseconds, giving the boat a speed of 2 millimeters per millisecond. It will then get 5 milliseconds to move, reaching a total distance of 10 millimeters.
				//   Hold the button for 3 milliseconds. After its remaining 4 milliseconds of travel time, the boat will have gone 12 millimeters.
				//   Hold the button for 4 milliseconds. After its remaining 3 milliseconds of travel time, the boat will have gone 12 millimeters.
				//   Hold the button for 5 milliseconds, causing the boat to travel a total of 10 millimeters.
				// The distances we see for each millisecond are [10 12 12 10], which is a palindrome.
				// Once we reach the center of the palindrome, then we just need to double the number of ways we can beat the record and break out of the loop.
				if distanceTravelled <= prevDistance && prevDistance != 0 {
					if distanceTravelled == prevDistance {
						waysToBeat *= 2
					} else {
						waysToBeat = (waysToBeat * 2) - 1
					}
					break
				} else {
					waysToBeat++
				}
				prevDistance = distanceTravelled
			}
		}
		recordBreakingWays = append(recordBreakingWays, waysToBeat)
	}

	return recordBreakingWays
}
