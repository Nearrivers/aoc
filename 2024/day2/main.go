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

func getGameDiceCountByColor(game, color string) int {
	diceMatch := regexp.MustCompile(`([0-9]+) `+color).FindAllStringSubmatch(game, -1)
	if len(diceMatch) > 0 {
		minCount := -9999
		for _, dice := range diceMatch {
			diceCount, _ := strconv.Atoi(dice[1])
			if minCount < diceCount {
				minCount = diceCount
			}
		}

		return minCount
	}

	return 1
}

func isGamePossible(game string, maxDices MaxDices) bool {
	sets := strings.Split(game, ";")

	for _, set := range sets {
		blueDicesCount := getGameSetDiceCountByColor(set, "blue")
		greenDicesCount := getGameSetDiceCountByColor(set, "green")
		redDicesCount := getGameSetDiceCountByColor(set, "red")

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

func getGamePower(game string) int {
	blueDicesMin := getGameDiceCountByColor(game, "blue")
	redDicesMin := getGameDiceCountByColor(game, "red")
	greenDicesMin := getGameDiceCountByColor(game, "green")
	return blueDicesMin * redDicesMin * greenDicesMin
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	total := 0
	totalPower := 0
	diceGames := strings.Split(fileContent, "\r\n")
	for _, game := range diceGames {
		id := getGameId(game)

		if isGamePossible(game, MaxDices{
			red:   12,
			blue:  14,
			green: 13,
		}) {
			total += id
		}

		power := getGamePower(game)
		totalPower += power
	}

	fmt.Println(total)
	fmt.Println(totalPower)
}
