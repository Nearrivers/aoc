package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxBlinkCount = 25

func blink(stones []int64, count int) []int64 {
	var newStones []int64	
	if count > maxBlinkCount {
		return stones
	}

	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

			strStone := strconv.FormatInt(stone, 10)
		if len(strStone) % 2 == 0 {
			h := len(strStone) / 2
			firstNewStone, _ := strconv.ParseInt(strStone[:int(h)], 10, 64)
			secondNewStone, _ := strconv.ParseInt(strStone[int(h):], 10, 64)

			newStones = append(newStones, []int64{firstNewStone, secondNewStone}...)
			continue
		}

		newStones = append(newStones, stone*2024)
	}

	count++
	return blink(newStones, count)
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	puzzle := string(b)

	strNb := strings.Fields(puzzle)
	numbers := make([]int64, 0)

	for _, s := range strNb {
		i, _ := strconv.ParseInt(s, 10, 64)
		numbers = append(numbers, i)
	}

	count := 1
	stones := blink(numbers, count)
	fmt.Println(len(stones))
}