package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getGameId(game string) int {
	re := regexp.MustCompile(`\d`)
	id, _ := strconv.Atoi(string(re.Find([]byte(game))))
	return id
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	diceGames := strings.Split(fileContent, "\r\n")
	for _, game := range diceGames {
		id := getGameId(game)
		fmt.Printf("%d\n", id)
	}
}
