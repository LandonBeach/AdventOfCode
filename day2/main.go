package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Round []Draw
	Max   map[string]int
}

type Draw struct {
	Colors map[string]int
}

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var IdTotal int = 0
	var powerTotal int = 0

	for scanner.Scan() {
		game, err := parseGameLine(scanner.Text())
		if err != nil {
			panic(err)
		}

		if checkValid(game) {
			IdTotal += game.Id
		}

		powerSubTotal := 1
		for _, numberBlocks := range game.Max {
			powerSubTotal *= numberBlocks
		}

		powerTotal += powerSubTotal
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	fmt.Println("ID total for valid games: ", IdTotal)
	fmt.Println("Power total for valid games: ", powerTotal)
}

// Parses each line from the game to get number of blocks drawn and their color.
func parseGameLine(line string) (Game, error) {
	var game Game

	// Split the Game ID and the rounds
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return game, fmt.Errorf("invalid input format: %s", line)
	}

	// Parse the Game ID
	gameId, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		return game, fmt.Errorf("error parsing game ID: %v", err)
	}
	game.Id = gameId

	game.Max = make(map[string]int)
	game.Max["red"] = 0
	game.Max["green"] = 0
	game.Max["blue"] = 0

	// Split out the rounds to parse each draw from the bag
	for _, roundStr := range strings.Split(parts[1], ";") {
		draw := Draw{
			Colors: make(map[string]int),
		}

		// Parse the number of blocks and the color
		for _, drawStr := range strings.Split(roundStr, ",") {
			part := strings.Split(strings.TrimSpace(drawStr), " ")
			color := strings.TrimSpace(part[1])

			numberBlocks, err := strconv.Atoi(strings.TrimSpace(part[0]))
			if err != nil {
				return game, fmt.Errorf("error parsing the draw from a round: %v", err)
			}

			draw.Colors[color] = numberBlocks

			if numberBlocks > game.Max[color] {
				game.Max[color] = numberBlocks
			}
		}

		game.Round = append(game.Round, draw)
	}

	return game, nil
}

// A game is "valid" if there are less than
// 12 reds, 13 blues, and 14 blues from each draw from the bag
func checkValid(game Game) bool {
	for _, draws := range game.Round {
		for color, value := range draws.Colors {
			switch color {
			case "red":
				if value > 12 {
					return false
				}
			case "green":
				if value > 13 {
					return false
				}
			case "blue":
				if value > 14 {
					return false
				}
			}
		}
	}
	return true
}
