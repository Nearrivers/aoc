package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	zero trailChar = '0'
	one trailChar = '1'
	two trailChar = '2'
	three trailChar = '3'
	four trailChar = '4'
	five trailChar = '5'
	six trailChar = '6'
	seven trailChar = '7'
	eight trailChar = '8'
	nine trailChar = '9'
)

type trailChar rune

type Trails [][]trailChar

var nextChar = map[trailChar]trailChar{
	one: two,
	two: three,
	three: four,
	four: five,
	five: six,
	six: seven,
	seven: eight,
	eight: nine,
}

func (t Trails) String() string {
	var out string

	for _, line := range t {
		out+= string(line)+"\n"
	}
	return out
}

func (t Trails) findTrail(i, j int, char trailChar) int {
	score := 0

	if char == nine {
		score++
		return score
	}

	adjacentNumber := make([]trailChar, 4)

	if i < len(t) -1 {
		adjacentNumber[0] = t[i][j]
	}

	if i
}

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	totalScore := 0
	t := make(Trails, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t = append(t, []trailChar(scanner.Text()))
	}

	fmt.Println(t)

	for i := range t {
		for j, tc := range t[i] {
			if tc == '0' {
				totalScore += t.findTrail(i, j, one)
			}
		}
	}

	fmt.Println(totalScore)
}