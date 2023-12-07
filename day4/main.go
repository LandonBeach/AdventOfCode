package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Card struct {
	WinningSet  map[string]int // Not needed, but good to know
	NumbersHave []string       // Not needed, but good to know
	MatchCount  int
	Instances   int
}

func main() {
	file := utils.OpenFile("input.txt")
	defer file.Close()

	var allCards []Card
	var totalValue int = 0
	scanner := bufio.NewScanner(file)

	// Regexp pattern to extract digits from each line/string
	digitPattern := `\b\d+\b`
	re := regexp.MustCompile(digitPattern)

	for scanner.Scan() {
		var matchCount float64 = 0

		// Card %d: <winning_numbers> | <given_numbers>
		scratchCard := strings.Split(scanner.Text(), ": ")
		cardNumbers := strings.Split(scratchCard[1], " | ")

		// Extract the winning numbers and the given numbers from the string
		winningNumbers := re.FindAllString(cardNumbers[0], -1)
		numbersHave := re.FindAllString(cardNumbers[1], -1)

		// Initialize the map/set containing the winning numbers
		//   Key: the winning number
		//   Value: number of occurrences within the given numbers
		winningSet := make(map[string]int)
		for _, num := range winningNumbers {
			winningSet[num] = 0
		}

		// Compare the number given/have with the winning numbers
		// Increase the number of occurrences if there is a match
		for _, num := range numbersHave {
			if _, exists := winningSet[num]; exists {
				winningSet[num]++
			}
		}

		// Get the total number of matches (number of occurrences excluded)
		for _, value := range winningSet {
			if value >= 1 {
				matchCount++
			}
		}

		// Calculate the scratch card points and add them all together
		if matchCount > 0 {
			totalValue += int(math.Pow(2.0, matchCount-1))
		}

		// Save the card for later...
		allCards = append(allCards, Card{
			WinningSet:  winningSet,
			NumbersHave: numbersHave,
			MatchCount:  int(matchCount),
			Instances:   1,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		panic(err)
	}

	// You win copies of the scratchcards below the winning card equal to the number of matches.
	// Copies win additional copies based on the number of matches.
	for cardId, card := range allCards {
		nextCardId := cardId + 1

		for i := 0; i < card.MatchCount; i++ {
			if nextCardId+i < len(allCards) {
				allCards[nextCardId+i].Instances += card.Instances
			}
		}
	}

	// Add up the total number of instances (original and copies)
	var totalCardInstances int = 0
	for _, card := range allCards {
		totalCardInstances += card.Instances
	}

	fmt.Println("Total of Scratch Card Points:", totalValue)
	fmt.Println("Total number of Scratch Card instances:", totalCardInstances)
}
