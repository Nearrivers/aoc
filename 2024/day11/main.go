package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memo map[[2]int64]int64

var memo Memo

func blink(remainingBlinks int64, stone int64) int64 {
	if remainingBlinks == 0 {
		return 1
	}

	if v, ok := memo[[2]int64{remainingBlinks, stone}]; ok {
		return v
	}

	if stone == 0 {
		result := blink(remainingBlinks-1, 1)
		memo[[2]int64{remainingBlinks, stone}] = result
		return result
	}

	strStone := strconv.FormatInt(stone, 10)
	if len(strStone)%2 == 0 {
		h := len(strStone) / 2
		firstNewStone, _ := strconv.ParseInt(strStone[:int(h)], 10, 64)
		secondNewStone, _ := strconv.ParseInt(strStone[int(h):], 10, 64)
		result := blink(remainingBlinks-1, firstNewStone) + blink(remainingBlinks-1, secondNewStone)
		memo[[2]int64{remainingBlinks, stone}] = result
		return result
	}

	result := blink(remainingBlinks-1, stone*2024)
	memo[[2]int64{remainingBlinks, stone}] = result
	return result
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	puzzle := string(b)
	strNb := strings.Fields(puzzle)
	var total int64

	memo = make(Memo)
	for _, s := range strNb {
		i, _ := strconv.ParseInt(s, 10, 64)
		total += blink(75, i)
	}

	fmt.Println(total)
}
