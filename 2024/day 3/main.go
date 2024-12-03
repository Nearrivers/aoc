package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Nearrivers/2024-day3-aoc/lexer"
	"github.com/Nearrivers/2024-day3-aoc/parser"
)

func main() {
	f, err := os.Open("realInput.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	totalResult := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)
		totalResult += p.ParseLine()
	}

	fmt.Println(totalResult)
}
