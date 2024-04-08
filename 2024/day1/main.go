package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MapLike struct {
	index int
	value int
}

func getTokenIndexes(calibration string) (int, int) {
	tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
	}

	var tokenIndexes []MapLike

	for token := range tokens {
		i := strings.Index(calibration, token)

		if i > -1 {
			tokenIndexes = append(tokenIndexes, MapLike{index: i, value: tokens[token]})
		}

		i = strings.LastIndex(calibration, token)
		if i > -1 {
			tokenIndexes = append(tokenIndexes, MapLike{index: i, value: tokens[token]})
		}
	}

	sort.Slice(tokenIndexes, func(i, j int) bool {
		return tokenIndexes[i].index < tokenIndexes[j].index
	})

	// for _, t := range tokenIndexes {
	// 	fmt.Printf("%s: index: %d, value: %d \n", calibration, t.index, t.value)
	// }

	return tokenIndexes[0].value, tokenIndexes[len(tokenIndexes)-1].value
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	var total int

	amendedCalibrationValues := strings.Split(fileContent, "\r\n")

	for _, amendCal := range amendedCalibrationValues {
		firstDigit, secondDigit := getTokenIndexes(amendCal)
		fmt.Printf("%s: %d - %d \n", amendCal, firstDigit, secondDigit)
		numberString := fmt.Sprintf("%d%d", firstDigit, secondDigit)
		number, _ := strconv.Atoi(numberString)
		total += number
	}

	fmt.Println(total)
}
