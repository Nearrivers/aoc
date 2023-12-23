package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	var total int

	amendedCalibrationValues := strings.Split(fileContent, "\r\n")

	for _, amendCal := range amendedCalibrationValues {
		var firstDigit, secondDigit string

		for _, runeChar := range amendCal {
			char := string(runeChar)

			_, err := strconv.Atoi(char)
			if err == nil {
				firstDigit = char
				break
			}
		}

		for i := len(amendCal) - 1; i >= 0; i-- {
			char := string(amendCal[i])

			_, err := strconv.Atoi(char)
			if err == nil {
				secondDigit = char
				break
			}
		}

		numberString := firstDigit + secondDigit

		number, _ := strconv.Atoi(numberString)

		total += number
	}

	fmt.Println(total)
}
