package main

import (
	"bufio"
	"fmt"
	"os"
)

type Puzzle struct {
	input [][]rune
	count int
}

func (p *Puzzle) String() string {
	var printedString string

	for _, line := range p.input {
		printedString += fmt.Sprintf("%b\n", line)
	}

	return printedString
}

func NewPuzzle() *Puzzle {
	file, err := os.Open("realInput.Txt")
	if err != nil {
		panic(err)
	}

	matrix := make([][]rune, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newLine := make([]rune, 0)

		for _, c := range line {
			newLine = append(newLine, c)
		}

		matrix = append(matrix, newLine)
	}

	return &Puzzle{
		input: matrix,
		count: 0,
	}
}

func main() {
	p := NewPuzzle()

	for i, line := range p.input {
		canReadUp := i >= 1
		canReadDown := i <= len(p.input)-2
		for j, letter := range line {
			canReadBackward := j >= 1
			canReadForward := j <= len(line)-2
			if letter == 'A' && canReadDown && canReadUp && canReadBackward && canReadForward {
				upLeft := p.input[i-1][j-1]
				upRight := p.input[i-1][j+1]
				downLeft := p.input[i+1][j-1]
				downRight := p.input[i+1][j+1]

				isFirstDiagonalGood := false
				isSecondDiagonalGood := false

				if (upLeft == 'M' && downRight == 'S') || (upLeft == 'S' && downRight == 'M') {
					isFirstDiagonalGood = true
				} else {
					continue
				}

				if (upRight == 'M' && downLeft == 'S') || (upRight == 'S' && downLeft == 'M') {
					isSecondDiagonalGood = true
				} else {
					continue
				}

				if isFirstDiagonalGood && isSecondDiagonalGood {
					p.count++
				}
			}
		}
	}

	fmt.Println(p.count)
}
