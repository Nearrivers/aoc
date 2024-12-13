package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

// I could have never found the answer alone. The answer was found following and understanding
// the wonderful video from HyperNeutrino at https://www.youtube.com/watch?v=-5J-DAsWuJc
func main() {
	var re = regexp.MustCompile(`\d+`)

	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strNumbers := re.FindAllString(string(b), -1)

	machines := make([][6]float64, 0)
	for chunk := range slices.Chunk(strNumbers, 6) {
		c := [6]float64{0, 0, 0, 0, 0, 0}
		for i := range chunk {
			n, _ := strconv.ParseFloat(chunk[i], 64)
			c[i] = n
		}

		machines = append(machines, c)
	}

	totalChips := 0
	for _, m := range machines {
		ax, ay, bx, by, px, py := m[0], m[1], m[2], m[3], m[4], m[5]
		px += 10000000000000
		py += 10000000000000
		nbA := (px*by - py*bx) / (ax*by - ay*bx)
		nbB := (px - ax*nbA) / bx
		if math.Mod(nbA, 1.0) == 0 && math.Mod(nbB, 1.0) == 0 {
			totalChips += int(nbA)*3 + int(nbB)
		}
	}

	fmt.Println(totalChips)
}
