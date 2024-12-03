package main

import (
	"fmt"
	"os"

	"github.com/Nearrivers/2024-day3-aoc/lexer"
	"github.com/Nearrivers/2024-day3-aoc/parser"
)

func main() {
	b, err := os.ReadFile("realInput.txt")
	if err != nil {
		panic(err)
	}

	totalResult := 0
	line := string(b)
	l := lexer.NewLexer(line)
	p := parser.NewParser(l)
	totalResult += p.ParseLine()

	fmt.Println(totalResult)
}
