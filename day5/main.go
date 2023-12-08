package main

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

type Entry struct {
	Seed        int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type Mapping struct {
	Destination int
	Source      int
	Range       int
}

// Initialize a new instance of a Mapping
func NewMapping(src int, dest int, mapRange int) *Mapping {
	return &Mapping{
		Destination: src,
		Source:      dest,
		Range:       mapRange,
	}
}

type Almanac struct {
	Maps map[string][]Mapping
}

// Initialize a new instance of an Almanac
func NewAlmanac() *Almanac {
	return &Almanac{
		Maps: make(map[string][]Mapping),
	}
}

func (al *Almanac) AddMapping(name string, mapping *Mapping) {
	al.Maps[name] = append(al.Maps[name], *mapping)
}

func (al Almanac) find(mapName string, src int, dest *int) {
	for _, mapping := range al.Maps[mapName] {
		if src >= mapping.Source && src < (mapping.Source+mapping.Range) {
			*dest = mapping.Destination + (src - mapping.Source)
			return
		}
	}

	*dest = src
}

func (al Almanac) FindEntryBySeed(seed int) Entry {
	entry := Entry{
		Seed: seed,
	}

	al.find("seed-to-soil", entry.Seed, &entry.Soil)
	al.find("soil-to-fertilizer", entry.Soil, &entry.Fertilizer)
	al.find("fertilizer-to-water", entry.Fertilizer, &entry.Water)
	al.find("water-to-light", entry.Water, &entry.Light)
	al.find("light-to-temperature", entry.Light, &entry.Temperature)
	al.find("temperature-to-humidity", entry.Temperature, &entry.Humidity)
	al.find("humidity-to-location", entry.Humidity, &entry.Location)

	fmt.Println("Entry:", entry)

	return entry
}

func main() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	file := strings.Split(string(fileImport), "\n")

	// Get seed values from the first line of the input file
	var seedValues []int
	for _, seed := range strings.Split(file[0], " ")[1:] {
		seedValues = append(seedValues, utils.StringToInt(seed))
	}

	almanac := NewAlmanac()

	// Name of each map according to the input file
	mapNames := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	// Extract all values for each map (under the map name in input file) and add them to the almanac
	for _, name := range mapNames {
		for _, mapping := range extractMap(file, name) {
			newMap := NewMapping(mapping[0], mapping[1], mapping[2])
			almanac.AddMapping(name, newMap)
		}
	}

	// Find the closest/lowest location number that corresponds to any of the initial seeds
	closestLocation := math.MaxInt
	for _, seed := range seedValues {
		entry := almanac.FindEntryBySeed(seed)
		if entry.Location < closestLocation {
			closestLocation = entry.Location
		}
	}

	fmt.Println("Closest Location:", closestLocation)
}

// Extract a map from the given input file/string
// A map starts with a name and its values are all the lines underneath it
// There is an empty/blank line separating each map.
func extractMap(file []string, mapName string) [][]int {
	var foundMap bool
	var mapping [][]int
	numberPattern := `\b(\d+)\b`
	re := regexp.MustCompile(numberPattern)

	for _, line := range file {
		// Find the map by name
		if strings.Contains(line, mapName) {
			foundMap = true
			continue
		}

		// Extract all the values underneath it
		if foundMap {
			if entry := re.FindAllString(line, -1); entry != nil {
				mapping = append(mapping, utils.StringsToInts(entry))
			} else {
				foundMap = false
				break
			}
		}
	}

	return mapping
}
