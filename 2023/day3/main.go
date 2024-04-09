package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func populateArray(lines []string) [][]rune {
	schematic := make([][]rune, len(lines))
	for i, line := range lines {
		schematic[i] = make([]rune, 0, len(line))
		for _, char := range line {
			schematic[i] = append(schematic[i], char)
		}
	}

	return schematic
}

func getNumber(j int, line string) (int, int) {
	numberString, offset, number := "", 0, 0

	for j+offset < len(line) {
		num, err := strconv.Atoi(string(line[j+offset]))
		if err != nil {
			break
		}

		numberString += fmt.Sprint(num)
		offset++
	}

	number, _ = strconv.Atoi(numberString)

	return number, len(numberString)
}

// ***** <-- On vérifie si le symbole est sur la ligne du dessus et sur les diagonales haute gauches et droites
//
// -463-
//
// *****
func isTheSymbolAbove(i, j int, numberLength int, schematic [][]rune) bool {
	// Si i == 0 alors nous sommes sur la première ligne, pas de ligne du dessus
	if i == 0 {
		return false
	}

	k := 0

	// Si j == 0 alors nous sommes à gauche donc pas de diagonale supérieure gauche
	if j == 0 {
		k = j
	} else {
		k = j - 1
	}

	charToSearch := 0
	// Si j == len(schematic[i]) - 1 alors nous sommes à droite donc pas de diagonale supérieure droite
	if j+numberLength == len(schematic[i]) {
		charToSearch = numberLength + 1
	} else {
		charToSearch = numberLength + 2
	}

	for l := 0; l < charToSearch; l++ {
		_, err := strconv.Atoi(string(schematic[i-1][k]))
		if err != nil && string(schematic[i-1][k]) != "." {
			return true
		}

		k++
	}

	return false
}

func isTheSymbolBelow(i, j int, numberLength int, schematic [][]rune) bool {
	// Si i == len(schematic) - 1 alors nous sommes sur la dernière ligne, pas de ligne du dessous
	if i == len(schematic)-1 {
		return false
	}

	k := 0

	// Si j == 0 alors nous sommes à gauche donc pas de diagonale inférieure gauche
	if j == 0 {
		k = j
	} else {
		k = j - 1
	}

	charToSearch := 0
	// Si j == len(schematic[i]) - 1 alors nous sommes à droite donc pas de diagonale inférieure droite
	if j+numberLength == len(schematic[i]) {
		charToSearch = numberLength + 1
	} else {
		charToSearch = numberLength + 2
	}

	for l := 0; l < charToSearch; l++ {
		_, err := strconv.Atoi(string(schematic[i+1][k]))
		if err != nil && string(schematic[i+1][k]) != "." {
			return true
		}

		k++
	}

	return false
}

func isTheSymbolToTheRight(i, j, numberLength int, schematic [][]rune) bool {
	if j+numberLength == len(schematic[i]) {
		return false
	}

	_, err := strconv.Atoi(string(schematic[i][j+numberLength]))
	if err != nil && string(schematic[i][j+numberLength]) != "." {
		return true
	}

	return false
}

func isTheSymbolToTheLeft(i, j int, schematic [][]rune) bool {
	if j == 0 {
		return false
	}

	_, err := strconv.Atoi(string(schematic[i][j-1]))
	if err != nil && string(schematic[i][j-1]) != "." {
		return true
	}

	return false
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	total := 0
	schematicLines := strings.Split(fileContent, "\r\n")
	schematic := populateArray(schematicLines)

	for i, line := range schematic {
		skipNext := 0
		for j, char := range schematic[i] {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			if skipNext > 0 {
				skipNext--
				continue
			}

			number, numberLength := getNumber(j, string(line))
			skipNext = numberLength - 1

			isAbove := isTheSymbolAbove(i, j, numberLength, schematic)
			isBelow := isTheSymbolBelow(i, j, numberLength, schematic)
			isRight := isTheSymbolToTheRight(i, j, numberLength, schematic)
			isLeft := isTheSymbolToTheLeft(i, j, schematic)

			if isAbove || isBelow || isRight || isLeft {
				total += number
			}
		}
	}

	fmt.Println(total)
}
