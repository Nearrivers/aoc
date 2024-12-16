package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	WIDTH  int64 = 101
	HEIGHT int64 = 103
)

type Grid [][]int

func (g Grid) String() string {
	out := ""
	for i := range g {
		for j := range g[i] {
			out += fmt.Sprintf("%d", g[i][j])
		}
		out += "\n"
	}
	return out
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`-?\d+`)
	robots := make([][4]int64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := re.FindAllString(scanner.Text(), -1)

		newRobots := [4]int64{0, 0, 0, 0}
		for i := range s {
			n, _ := strconv.ParseInt(s[i], 10, 64)
			newRobots[i] = n
		}
		robots = append(robots, newRobots)
	}

	finalCoords := make([][2]int64, 0)

	for _, robot := range robots {
		px, py, vx, vy := robot[0], robot[1], robot[2], robot[3]
		cx := (px + vx*100) % WIDTH
		cy := (py + vy*100) % HEIGHT
		if cx < 0 {
			cx = WIDTH + cx
		}

		if cy < 0 {
			cy = HEIGHT + cy
		}
		// Modulo largeur car les robots se téléportent de l'autre côté de la grille une fois qu'ils la quitte
		finalCoords = append(finalCoords, [2]int64{cx, cy})
	}

	grid := make(Grid, 0)
	for i := range HEIGHT {
		line := make([]int, WIDTH)
		for j := range WIDTH {
			line[j] = 0
			for _, coord := range finalCoords {
				if coord[0] == j && coord[1] == i {
					line[j]++
				}
			}
		}
		grid = append(grid, line)
	}

	Xsep := WIDTH / 2
	Ysep := HEIGHT / 2

	quadrantsSums := []int{0, 0, 0, 0}
	for i := range grid {
		if i == int(Ysep) {
			continue
		}
		for j := range grid[i] {
			if j == int(Xsep) {
				continue
			}
			if i < int(Ysep) && j < int(Xsep) {
				quadrantsSums[0] += grid[i][j]
				continue
			}

			if i < int(Ysep) && j > int(Xsep) {
				quadrantsSums[1] += grid[i][j]
				continue
			}

			if i > int(Ysep) && j < int(Xsep) {
				quadrantsSums[2] += grid[i][j]
				continue
			}

			if i > int(Ysep) && j > int(Xsep) {
				quadrantsSums[3] += grid[i][j]
			}
		}
	}

	total := 1
	for i := range quadrantsSums {
		total *= quadrantsSums[i]
	}
	fmt.Println(total)
}
