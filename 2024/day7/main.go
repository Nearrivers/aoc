package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// I did not find the answer alone. The answer was found following
// and understanding the code found in this repository https://github.com/pin2t/aoc2024/blob/master/07.go

func addition(calibration, mesure int) int {
	return calibration + mesure
}

func multiplication(calibration, mesure int) int {
	return calibration * mesure
}

func concat(calibration, mesure int) int {
	strCal := strconv.FormatInt(int64(calibration), 10)
	strMes := strconv.FormatInt(int64(mesure), 10)

	strResult := strCal + strMes

	result, _ := strconv.ParseInt(strResult, 10, 64)
	return int(result)
}

func match(numbers []int, index, calibration int) bool {
	if index == len(numbers) - 1 {
		return numbers[0] == addition(calibration, numbers[index]) ||
		numbers[0] == multiplication(calibration, numbers[index]) ||
		numbers[0] == concat(calibration, numbers[index])
	}

	return match(numbers, index+1, addition(calibration, numbers[index])) ||
		match(numbers, index+1, multiplication(calibration, numbers[index])) ||
		match(numbers, index+1, concat(calibration, numbers[index]))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	reg := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var numbers []int
		strNumbers := reg.FindAllString(scanner.Text(), -1)
		for _, s := range strNumbers {
			n, _ := strconv.Atoi(s)
			numbers = append(numbers, n)
		}

		if match(numbers, 1, 0) {
			total += numbers[0]
		}
	}

	fmt.Println(total)
}