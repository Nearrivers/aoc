package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

const (
	DIMENSIONS = 71
	BYTECOUNT  = 1024
)

type Vertex struct {
	row, col, distance int
}

type MinHeap []Vertex

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Vertex))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Grid [][]rune

func (g Grid) String() string {
	var out string

	for i := range g {
		out += string(g[i]) + "\n"
	}
	return out
}

type Move struct {
	row, col int
}

type Visited map[[2]int]bool

func (v Visited) String() string {
	var out string

	for k := range v {
		out += fmt.Sprintf("[%d, %d]\n", k[0], k[1])
	}
	return out
}

func main() {
	grid := make(Grid, DIMENSIONS)

	for i := range grid {
		grid[i] = make([]rune, DIMENSIONS)
		for j := range grid[0] {
			grid[i][j] = '.'
		}
	}

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	count := 0
	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(`\d+`)

	for scanner.Scan() && count < BYTECOUNT {
		count++
		s := reg.FindAllString(scanner.Text(), -1)

		n1, _ := strconv.ParseInt(s[0], 10, 64)
		n2, _ := strconv.ParseInt(s[1], 10, 64)
		grid[n2][n1] = '#'
	}

	var priorityQueue MinHeap

	heap.Push(&priorityQueue, Vertex{
		col:      0,
		row:      0,
		distance: 0,
	})

	visited := make(Visited, 0)
	visited[[2]int{0, 0}] = true
	distances := make([]int, 0)

outer:
	for priorityQueue.Len() > 0 {
		v := heap.Pop(&priorityQueue).(Vertex)

		possibleMoves := []Move{
			{row: v.row + 1, col: v.col},
			{row: v.row - 1, col: v.col},
			{row: v.row, col: v.col + 1},
			{row: v.row, col: v.col - 1},
		}

		grid[v.row][v.col] = 'x'

		for _, move := range possibleMoves {
			if move.row < 0 || move.col < 0 || move.row == DIMENSIONS || move.col == DIMENSIONS {
				continue
			}

			if grid[move.row][move.col] == '#' {
				continue
			}

			mn := [2]int{move.row, move.col}
			if _, ok := visited[mn]; ok {
				continue
			}

			distances = append(distances, v.distance)

			if move.col == DIMENSIONS-1 && move.row == DIMENSIONS-1 {
				fmt.Println(v.distance + 1)
				break outer
			}

			visited[mn] = true
			heap.Push(&priorityQueue, Vertex{
				row:      move.row,
				col:      move.col,
				distance: v.distance + 1,
			})
		}
	}
	// fmt.Println(visited)
	distances = slices.Compact(distances)
	fmt.Println(len(distances))
}
