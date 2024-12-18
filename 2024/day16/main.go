package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"math"
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
	// Map du dernier meilleur chemin emprunté pour chaque sommet
	lastPath := make(map[[4]int]map[[4]int]bool, 0)
	endMoves := make(map[[4]int]bool, 0)
	minDistance := math.MaxInt

	for len(priorityQueue) > 0 {
		v := heap.Pop(&priorityQueue).(Vertex)

		vn := [4]int{v.row, v.col, v.dr, v.dc}
		// Initialisation distance première visite à +Inf
		_, ok := visited[vn]
		if !ok {
			visited[vn] = math.MaxInt
		}

		// Si on a déjà trouvé un meilleur chemin, on passe
		if v.distance > visited[vn] {
			continue
		}

		if grid[v.row][v.col] == 'E' {
			// Si on est arrivé au bout avec un pire chemin que le meilleur trouvé, alors on quitte.
			// Etant donné que l'on utilise une minHeap, si on trouve un chemin avec un total de distance supérieur, alors
			// cela veut dire que le ou les meilleurs chemins on été trouvés et que l'on peut arrêter l'exécution
			if v.distance > minDistance {
				break
			}
			minDistance = v.distance
			endMoves[vn] = true
		}

		possibleMoves := []Move{
			{cost: v.distance + 1, row: v.row + v.dr, col: v.col + v.dc, dr: v.dr, dc: v.dc},
			{cost: v.distance + 1000, row: v.row, col: v.col, dr: v.dc, dc: -v.dr},
			{cost: v.distance + 1000, row: v.row, col: v.col, dr: -v.dc, dc: v.dr},
		}

		for _, move := range possibleMoves {
			n := [4]int{move.row, move.col, move.dr, move.dc}

			lowestDistance, ok := visited[n]
			if !ok {
				lowestDistance = math.MaxInt
			}

			if grid[move.row][move.col] == '#' || (lowestDistance < move.cost) {
				continue
			}

			if move.cost < lowestDistance {
				lastPath[n] = make(map[[4]int]bool)
				// Ajout du chemin
				visited[n] = move.cost
			}

			if grid[move.row][move.col] != 'E' {
				grid[move.row][move.col] = 'x'
			}

			lastPath[n][vn] = true

			heap.Push(&priorityQueue, Vertex{
				row:      move.row,
				col:      move.col,
				dr:       move.dr,
				dc:       move.dc,
				distance: move.cost,
			})
		}
	}

	queue := list.New()
	fmt.Println(endMoves)

	seen := make(map[[4]int]bool, 0)

	for k := range endMoves {
		queue.PushBack(lastPath[k])
		seen[k] = true
	}

	tilesCount := 0

	for queue.Len() > 0 {
		front := queue.Front()

		for k := range front.Value.(map[[4]int]bool) {
			if _, ok := seen[k]; ok {
				continue
			}
			fmt.Println(k)

			seen[k] = true
			queue.PushBack(lastPath[k])
			tilesCount++
		}
		queue.Remove(front)
	}

	uniqueCoor := make(map[[2]int]bool, 0)

	for k := range seen {
		uk := [2]int{k[0], k[1]}
		uniqueCoor[uk] = true
	}
	fmt.Println(len(uniqueCoor))
}
