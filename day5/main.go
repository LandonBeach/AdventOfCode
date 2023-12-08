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

type Almanac struct {
	Maps map[string][]Mapping
}

func NewAlmanac() *Almanac {
	return &Almanac{
		Maps: make(map[string][]Mapping),
	}
}

func (al *Almanac) AddMapping(name string, mapping Mapping) {
	al.Maps[name] = append(al.Maps[name], mapping)
}

func (al *Almanac) find(mapName string, from int) int {
	for _, mapping := range al.Maps[mapName] {
		if from >= mapping.Source && from < (mapping.Source+mapping.Range) {
			return mapping.Destination + (from - mapping.Source)
		}
	}

	return from
}

func (al *Almanac) FindSeedEntry(seed int) Entry {
	entry := Entry{
		Seed: seed,
	}

	entry.Soil = al.find("seed-to-soil", entry.Seed)
	entry.Fertilizer = al.find("soil-to-fertilizer", entry.Soil)
	entry.Water = al.find("fertilizer-to-water", entry.Fertilizer)
	entry.Light = al.find("water-to-light", entry.Water)
	entry.Temperature = al.find("light-to-temperature", entry.Light)
	entry.Humidity = al.find("temperature-to-humidity", entry.Temperature)
	entry.Location = al.find("humidity-to-location", entry.Humidity)

	fmt.Println("Entry:", entry)

	return entry
}

func main() {
	fileImport, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	file := strings.Split(string(fileImport), "\n")

	var seedValues []int
	for _, seed := range strings.Split(file[0], " ")[1:] {
		seedValues = append(seedValues, utils.StringToInt(seed))
	}

	almanac := NewAlmanac()

	mapNames := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	// Seed to soil map
	for _, name := range mapNames {
		for _, mapping := range extractMap(file, name) {
			almanac.AddMapping(name, createMapping(mapping))
		}
	}

	closestLocation := math.MaxInt
	for _, seed := range seedValues {
		entry := almanac.FindSeedEntry(seed)
		if entry.Location < closestLocation {
			closestLocation = entry.Location
		}
	}

	fmt.Println("Closest Location:", closestLocation)
}

func createMapping(mapping []int) Mapping {
	newMapping := Mapping{
		Destination: mapping[0],
		Source:      mapping[1],
		Range:       mapping[2],
	}

	return newMapping
}

func extractMap(file []string, mapName string) [][]int {
	var foundMap bool
	var mapping [][]int
	numberPattern := `\b(\d+)\b`
	re := regexp.MustCompile(numberPattern)

	for _, line := range file {
		if strings.Contains(line, mapName) {
			foundMap = true
			continue
		}

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
