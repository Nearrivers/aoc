package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MaxDices struct {
	red   int
	blue  int
	green int
}

func getGameId(game string) int {
	re := regexp.MustCompile(`([0-9]+)`)
	id, _ := strconv.Atoi(string(re.Find([]byte(game))))
	return id
}

func getGameSetDiceCountByColor(set, color string) int {
	diceMatch := regexp.MustCompile(`([0-9]+) ` + color).FindStringSubmatch(set)
	if len(diceMatch) > 0 {
		diceCount, _ := strconv.Atoi(diceMatch[1])
		return diceCount
	}

	return 0
}

func isGamePossible(game string, maxDices MaxDices) bool {
	sets := strings.Split(game, ";")

	for _, set := range sets {
		blueDicesCount := getGameSetDiceCountByColor(set, "blue")
		greenDicesCount := getGameSetDiceCountByColor(set, "green")
		redDicesCount := getGameSetDiceCountByColor(set, "red")

		fmt.Printf("set: %s || R: %d G: %d B: %d\n", set, redDicesCount, blueDicesCount, greenDicesCount)
		if blueDicesCount > maxDices.blue {
			return false
		}

		if greenDicesCount > maxDices.green {
			return false
		}

		if redDicesCount > maxDices.red {
			return false
		}
	}
	return true
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	total := 0
	diceGames := strings.Split(fileContent, "\r\n")
	for _, game := range diceGames {
		id := getGameId(game)

		if isGamePossible(game, MaxDices{
			blue:  14,
			red:   12,
			green: 13,
		}) {
			total += id
		}
	}

	fmt.Println(total)
}
