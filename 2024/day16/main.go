package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Grid [][]rune

func (g Grid) String() string {
	var out string
	for i := range g {
		out += string(g[i]) + "\n"
	}
	return out
}

type Vertex struct {
	row      int
	col      int
	dr       int
	dc       int
	distance int
}

var priorityQueue MinHeap

// On regarde les endroits où l'on a tourné et dans quelle direction on y est arrivé
var visited = make(map[[4]int]int, 0)

type Move struct {
	cost int
	row  int
	col  int
	dr   int
	dc   int
}

var grid = make(Grid, 0)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	deerRow, deerCol := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	file.Close()

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 'S' {
				deerCol = j
				deerRow = i
			}
		}
	}

	dr, dc := 0, 1
	heap.Push(&priorityQueue, Vertex{
		row:      deerRow,
		col:      deerCol,
		dr:       dr,
		dc:       dc,
		distance: 0,
	})

	visited[[4]int{deerRow, deerCol, dr, dc}] = 0

	for len(priorityQueue) > 0 {
		v := heap.Pop(&priorityQueue).(Vertex)
		visited[[4]int{v.row, v.col, v.dr, v.dc}] = v.distance

		if grid[v.row][v.col] == 'E' {
			// minDistance = v.distance
			fmt.Println(v.distance)
			break
		}

		possibleMoves := []Move{
			{cost: v.distance + 1, row: v.row + v.dr, col: v.col + v.dc, dr: v.dr, dc: v.dc},
			{cost: v.distance + 1000, row: v.row, col: v.col, dr: v.dc, dc: -v.dr},
			{cost: v.distance + 1000, row: v.row, col: v.col, dr: -v.dc, dc: v.dr},
		}

		for _, move := range possibleMoves {
			n := [4]int{move.row, move.col, move.dr, move.dc}

			vis, ok := visited[n]
			if grid[move.row][move.col] == '#' || (vis < v.distance && ok) {
				continue
			}

			heap.Push(&priorityQueue, Vertex{
				row:      move.row,
				col:      move.col,
				dr:       move.dr,
				dc:       move.dc,
				distance: move.cost,
			})
		}
	}
}
