package main

import (
	"bufio"
	"fmt"
	"os"
)

type trailChar rune

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

type coord struct {
	i, j int
}

type Trails [][]trailChar

var trailsEndCoord = []coord{}

var nextChar = map[trailChar]trailChar{
	zero: one,
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

func (t Trails) Print(i, j int) {
	var out string

	for k := range t {
		for l, char := range t[k] {
			if k == i && l == j {
				out += "x"
				continue
			}

			out += string(char)
		}
		out+= "\n"
	}

	fmt.Println(out)
}


func (t Trails) findTrail(i, j int, char trailChar) int {
	score := 0

	// t.Print(i, j)
	// fmt.Println(score)

	if char == nine {
		score++
		// if slices.IndexFunc(trailsEndCoord, func(c coord) bool {
		// 	return c.i == i && c.j == j
		// }) == - 1 {
		// 	score++
		// 	trailsEndCoord = append(trailsEndCoord, coord{i, j})
		// }
	}

	nc, ok := nextChar[char]
	if !ok {
		return score
	}

	if i < len(t) -1 && t[i+1][j] == nc {
		score += t.findTrail(i+1, j, nc)
	}

	if i > 0 && t[i-1][j] == nc {
		score += t.findTrail(i-1, j, nc)
	}

	if j < len(t[i]) - 1 && t[i][j+1] == nc {
		score += t.findTrail(i, j+1, nc)
	}

	if j > 0  && t[i][j-1] == nc {
		score += t.findTrail(i, j-1, nc)
	}

	return score
}

func main() {
	file, err := os.Open("input.txt")
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
				totalScore += t.findTrail(i, j, zero)
				trailsEndCoord = []coord{}
			}
		}
	}

	fmt.Println(totalScore)
}