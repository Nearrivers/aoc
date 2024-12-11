package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxBlinkCount = 75

func blink(stones []int64, count int, ch chan []int64) {
	var newStones []int64	
	if count > maxBlinkCount {
		ch <- stones
		return
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
	blink(newStones, count, ch)
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

	h := len(numbers) / 2
	firstCh := make(chan []int64)
	blink(numbers[:h], count, firstCh)

	secondCh := make(chan []int64)
	blink(numbers[h:], count, secondCh)

	var firstChunk []int64
	var secondChunk []int64
	for {
		select {
		case s := <- firstCh:
			firstChunk = s
			close(firstCh)
		case s := <- secondCh:
			secondChunk = s
			close(secondCh)
		}

		_, ok1 := <- firstCh
		_, ok2 := <- secondCh

		if !ok1 && !ok2 {
			break
		}
	}

	fmt.Println(len(firstChunk) + len(secondChunk))
}