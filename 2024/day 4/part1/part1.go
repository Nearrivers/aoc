package part1

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
)

var wordToSearch = []rune("XMAS")

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

func (p *Puzzle) readForward(i, j int) {
	line := p.input[i]

	piece := line[j : j+len(wordToSearch)]

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en avant")
		p.count++
	}
}

func (p *Puzzle) readBackward(i, j int) {
	line := p.input[i]

	reversedWord := []rune{'X', 'M', 'A', 'S'}
	slices.Reverse(reversedWord)
	piece := line[j-(len(wordToSearch)-1) : j+1]

	if reflect.DeepEqual(piece, reversedWord) {
		fmt.Println("+1 en arrière")
		p.count++
	}
}

func (p *Puzzle) readUp(i, j int) {
	var piece []rune

	index := i
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[index][j])
		index--
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en haut")
		p.count++
	}
}

func (p *Puzzle) readDown(i, j int) {
	var piece []rune

	index := i
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[index][j])
		index++
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en bas là")
		p.count++
	}
}

func (p *Puzzle) readUpLeft(i, j int) {
	var piece []rune

	rowIndex := i
	lineIndex := j
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[rowIndex][lineIndex])
		rowIndex--
		lineIndex--
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en haut à gauche")
		p.count++
	}
}

func (p *Puzzle) readUpRight(i, j int) {
	var piece []rune

	rowIndex := i
	lineIndex := j
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[rowIndex][lineIndex])
		rowIndex--
		lineIndex++
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en haut à droite")
		p.count++
	}
}

func (p *Puzzle) readDownLeft(i, j int) {
	var piece []rune

	rowIndex := i
	lineIndex := j
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[rowIndex][lineIndex])
		rowIndex++
		lineIndex--
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en bas à gauche")
		p.count++
	}
}

func (p *Puzzle) readDownRight(i, j int) {
	var piece []rune

	rowIndex := i
	lineIndex := j
	for len(piece) < len(wordToSearch) {
		piece = append(piece, p.input[rowIndex][lineIndex])
		rowIndex++
		lineIndex++
	}

	if reflect.DeepEqual(piece, wordToSearch) {
		fmt.Println("+1 en bas à droite")
		p.count++
	}
}

func main() {
	p := NewPuzzle()

	for i, line := range p.input {
		canReadUp := i >= len(wordToSearch)-1
		canReadDown := i <= (len(p.input) - len(wordToSearch))
		for j, letter := range line {
			if letter == 'X' {
				canReadBackward := j >= len(wordToSearch)-1
				canReadForward := j <= len(line)-len(wordToSearch)

				if canReadForward {
					p.readForward(i, j)
				}

				if canReadBackward {
					p.readBackward(i, j)
				}

				if canReadUp {
					p.readUp(i, j)

					if canReadForward {
						p.readUpRight(i, j)
					}

					if canReadBackward {
						p.readUpLeft(i, j)
					}
				}

				if canReadDown {
					p.readDown(i, j)

					if canReadForward {
						p.readDownRight(i, j)
					}

					if canReadBackward {
						p.readDownLeft(i, j)
					}
				}
			}
		}
	}

	fmt.Println(p.count)
}
