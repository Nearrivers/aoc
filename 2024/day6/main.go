package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid [][]rune

type VisitedSet map[[4]int]bool

var visited VisitedSet

func (g Grid) String() string {
	var out string
	for _, row := range g {
		out += "["
		for _, column := range row {
			out += string(column)
		}
		out += "]\n"
	}
	return out
}

func Loop(grid Grid, row, col int) bool {
	// Indice trouvé sur la vidéo https://youtu.be/2td0PZRKkpQ?si=GAlobfr1Smy_oz8D&t=301 (avec timestamps)
	deltaRow := -1
	deltaCol := 0

	visited = make(map[[4]int]bool)

	for {
		if row+deltaRow < 0 || row+deltaRow >= len(grid) || col+deltaCol < 0 || col+deltaCol >= len(grid[0]) {
			return false
		}
		visited[[4]int{row, col, deltaRow, deltaCol}] = true

		if grid[row+deltaRow][col+deltaCol] == '#' {
			// Tourner à droite
			deltaRow, deltaCol = deltaCol, -deltaRow
		} else {
			row += deltaRow
			col += deltaCol
		}

		if _, ok := visited[[4]int{row, col, deltaRow, deltaCol}]; ok {
			return true
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	grid := make(Grid, 0)

	scanner := bufio.NewScanner(f)
	count := -1
	soldierRow := 0
	soldierColumn := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		count++

		if strings.Contains(line, "^") {
			soldierColumn = strings.Index(line, "^")
			soldierRow = count
		}
	}

	obstructionsCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				continue
			}

			// On insère un obstacle
			grid[i][j] = '#'

			// On simule une ronde du garde avec cet obstacle
			if Loop(grid, soldierRow, soldierColumn) {
				obstructionsCount++
			}

			// On enlève le faux obstacle
			grid[i][j] = '.'
		}
	}

	fmt.Println(obstructionsCount)
}
