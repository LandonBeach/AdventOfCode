package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type EnginePart struct {
	Value       int
	OffsetStart int
	OffsetEnd   int
}

type Gear struct {
	Offset int
	ratio  []int
}

func (g *Gear) addRatio(value int) {
	if g.ratio == nil {
		g.ratio = make([]int, 0)
	}

	g.ratio = append(g.ratio, value)
}

// A ratio is calculated only when the gear
// is attached to exactly 2 engine parts
func (g *Gear) Ratio() int {
	if len(g.ratio) == 2 {
		return g.ratio[0] * g.ratio[1]
	}

	return 0
}

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	var prevLine string
	var enginePartSum int
	var gearRatioSum int
	var currentEnginePart []EnginePart
	var prevEnginePart []EnginePart
	var currentGears []Gear
	var prevGears []Gear

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentLine := scanner.Text()
		currentEnginePart, _ = findEngineParts(currentLine)
		currentGears = findGears(currentLine)

		// PART 1: Find engine parts, these are numbers that are adjacent to symbols
		// Check the previous iterations engine parts in the current line
		checkEnginePart(currentLine, &prevEnginePart, &enginePartSum)

		// Check for engine parts from this iteration in the previous line
		checkEnginePart(prevLine, &currentEnginePart, &enginePartSum)

		// Check for engine parts from this iteration in the current line
		checkEnginePart(currentLine, &currentEnginePart, &enginePartSum)

		// PART 2: Find gears ("*" symbol) that are attached to exactly 2 engine parts
		// Check the current engine parts for gears from the previous iteration
		findAttachedGears(&currentEnginePart, &prevGears)

		// Check the previous iterations engine parts for gears from the current iteration
		findAttachedGears(&prevEnginePart, &currentGears)

		// Check the current engine parts for gears from the current iteration
		findAttachedGears(&currentEnginePart, &currentGears)

		// Add the gear ratio from the previous iteration
		// Current iteration gears haven't checked the next line yet so they are incomplete and won't work.
		for _, gears := range prevGears {
			gearRatioSum += gears.Ratio()
		}

		// Create a new array by making a slice with the same length as gears
		prevGears = make([]Gear, len(currentGears))
		copy(prevGears, currentGears)

		prevLine = currentLine
		prevEnginePart = currentEnginePart
	}

	// Don't forget to add up the gear ratio from the last line
	for _, gear := range currentGears {
		gearRatioSum += gear.Ratio()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	fmt.Println("Sum of all part numbers:", enginePartSum)
	fmt.Println("Sum of all gear ratios:", gearRatioSum)
}

// Find all of the gears ("*" symbol) within a line/string
// The Offset is the index of the gear. By deafult the gear isn't attached to any engine parts so the ratio is empty
func findGears(line string) []Gear {
	var gears []Gear
	gearPattern := `\*`
	re := regexp.MustCompile(gearPattern)
	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		gear := Gear{
			Offset: match[0],
			ratio:  []int{},
		}
		gears = append(gears, gear)
	}

	return gears
}

// Checks to see if the gears are attached (adjacent) to any engine parts
func findAttachedGears(enginePart *[]EnginePart, gears *[]Gear) {
	for _, part := range *enginePart {
		beginning := part.OffsetStart - 1
		end := part.OffsetEnd

		for i := 0; i < len(*gears); i++ {
			if (*gears)[i].Offset >= beginning && (*gears)[i].Offset <= end {
				(*gears)[i].addRatio(part.Value)
			}
		}
	}
}

// Checks to see if it is a "valid" engine part by determining if there is a symbol adjacent to it.
func checkEnginePart(line string, enginePart *[]EnginePart, engineSum *int) {
	if line == "" {
		return
	}

	for _, part := range *enginePart {
		beginning := part.OffsetStart - 1
		end := part.OffsetEnd + 1

		if beginning < 0 {
			beginning++
		}

		if end >= len(line) {
			end--
		}

		// Check if the line contains any symbols. If true then add
		// the number value to the total.
		if strings.ContainsAny(line[beginning:end], "@#$%&*-+=/") {
			*engineSum += part.Value
		}
	}
}

// Find all engine parts within a line/string. An engine part is a group of digits surrounded by a period and/or symbol.
// The offset is the start and end index of the engine part within the line/string.
func findEngineParts(line string) ([]EnginePart, error) {
	var foundEngineParts []EnginePart

	enginePattern := `\b(\d+)\b`
	re := regexp.MustCompile(enginePattern)
	matches := re.FindAllString(line, -1)
	indexes := re.FindAllStringIndex(line, -1)

	for i := 0; i < len(matches); i++ {
		value, err := strconv.Atoi(matches[i])
		if err != nil {
			return foundEngineParts, fmt.Errorf("error parsing digit in match")
		}

		foundEngineParts = append(foundEngineParts, EnginePart{
			Value:       value,
			OffsetStart: indexes[i][0],
			OffsetEnd:   indexes[i][1],
		})
	}

	return foundEngineParts, nil
}
