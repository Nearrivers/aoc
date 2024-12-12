package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
)

type Garden [][]rune

var visited [][2]int
var garden = make(Garden, 0)

func FindArrangement(plant rune, i, j int) (int, int) {
	if slices.ContainsFunc(visited, func(v [2]int) bool {
		return reflect.DeepEqual(v, [2]int{i, j})
	}) {
		return 0, 0
	}

	visited = append(visited, [2]int{i, j})

	perimeter := 4
	area := 1
	var isRightSame bool
	var isDownSame bool
	var isLeftSame bool
	var isUpSame bool

	if j < len(garden[i])-1 && garden[i][j+1] == plant {
		isRightSame = true
		perimeter--
	}

	if j > 0 && garden[i][j-1] == plant {
		isLeftSame = true
		perimeter--
	}

	if i < len(garden)-1 && garden[i+1][j] == plant {
		isDownSame = true
		perimeter--
	}

	if i > 0 && garden[i-1][j] == plant {
		isUpSame = true
		perimeter--
	}

	if isDownSame {
		nextPerim, nextArea := FindArrangement(plant, i+1, j)
		perimeter += nextPerim
		area += nextArea
	}

	if isRightSame {
		nextPerim, nextArea := FindArrangement(plant, i, j+1)
		perimeter += nextPerim
		area += nextArea
	}

	if isUpSame {
		nextPerim, nextArea := FindArrangement(plant, i-1, j)
		perimeter += nextPerim
		area += nextArea
	}

	if isLeftSame {
		nextPerim, nextArea := FindArrangement(plant, i, j-1)
		perimeter += nextPerim
		area += nextArea
	}

	return perimeter, area
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, []rune(line))
	}

	totalPrice := 0

	for i := range garden {
		for j := range garden[i] {
			plant := garden[i][j]
			totalPerimeter, totalArea := FindArrangement(plant, i, j)

			totalPrice += totalPerimeter * totalArea
			if totalPerimeter == totalArea && totalPerimeter == 0 {
				continue
			}
			fmt.Printf("Plant: %s, périmètre: %d, aire: %d, total: %d\n", string(plant), totalPerimeter, totalArea, totalPerimeter*totalArea)
		}
	}

	fmt.Println(totalPrice)
}
