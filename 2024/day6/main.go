package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid [][]rune

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

func moveLeft(rowIndex, columnIndex int, grid Grid) int {
	count := 0
	if grid[rowIndex][columnIndex] != 'X' {
		grid[rowIndex][columnIndex] = 'X'
		count++
		fmt.Println(grid)
	}

	if columnIndex == 0 {
		return count
	}

	if grid[rowIndex][columnIndex-1] == '#' {
		count += moveUp(rowIndex, columnIndex, grid)
		return count
	}

	count += moveLeft(rowIndex, columnIndex-1, grid)
	return count
}

func moveDown(rowIndex, columnIndex int, grid Grid) int {
	count := 0
	if grid[rowIndex][columnIndex] != 'X' {
		grid[rowIndex][columnIndex] = 'X'
		count++
		fmt.Println(grid)
	}

	if rowIndex == len(grid)-1 {
		return count
	}

	if grid[rowIndex+1][columnIndex] == '#' {
		count += moveLeft(rowIndex, columnIndex, grid)
		return count
	}

	count += moveDown(rowIndex+1, columnIndex, grid)
	return count
}

func moveRight(rowIndex, columnIndex int, grid Grid) int {
	count := 0
	if grid[rowIndex][columnIndex] != 'X' {
		grid[rowIndex][columnIndex] = 'X'
		count++
		fmt.Println(grid)
	}

	if columnIndex == len(grid[rowIndex])-1 {
		return count
	}

	if grid[rowIndex][columnIndex+1] == '#' {
		count += moveDown(rowIndex, columnIndex, grid)
		return count
	}

	count += moveRight(rowIndex, columnIndex+1, grid)
	return count
}

func moveUp(rowIndex, columnIndex int, grid Grid) int {
	count := 0
	if grid[rowIndex][columnIndex] == '.' {
		grid[rowIndex][columnIndex] = '|'
		count++
		fmt.Println(grid)
	}

	if rowIndex == 0 {
		return count
	}

	if grid[rowIndex-1][columnIndex] == '#' {
		grid[rowIndex][columnIndex] = '+'
		count += moveRight(rowIndex, columnIndex, grid)
		return count
	}

	count += moveUp(rowIndex-1, columnIndex, grid)
	return count
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

	total := moveUp(soldierRow, soldierColumn, grid)

	fmt.Println(total)
}
