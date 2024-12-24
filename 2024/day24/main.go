package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Input map[string]int

func (i Input) checkKey(key string) bool {
	_, ok := i[key]
	return ok
}

var inputs = make(Input, 0)

type Wire struct {
	i1, i2, output string
}

func parseLine(line []string, ch chan string, mu *sync.Mutex) {
	i1 := line[0]
	i2 := line[2]
	op := line[1]
	output := line[4]

	for {
		mu.Lock()
		if inputs.checkKey(i1) && inputs.checkKey(i2) {
			switch op {
			case "AND":
				inputs[output] = inputs[i1] & inputs[i2]
			case "OR":
				inputs[output] = inputs[i1] | inputs[i2]
			case "XOR":
				inputs[output] = inputs[i1] ^ inputs[i2]
			}	
			break
		}
		mu.Unlock()
	}
}

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	isSecondPart := false
	strCh := make(chan string)
	mu := &sync.Mutex{}

	wg := sync.WaitGroup{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isSecondPart = true
			continue
		}

		if isSecondPart {
			wg.Add(1)
			go func ()  {
				parseLine(strings.Fields(line), strCh, mu)
				wg.Done()
			}()
			continue
		}

		inp := strings.Split(line, ": ")
		val, _ := strconv.Atoi(inp[1])
		inputs[inp[0]] = val
	}

	fmt.Println(inputs)
}