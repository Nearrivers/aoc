package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isUpdateOrdered(rules [][]int, update []int) bool {
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) &&
			slices.Contains(update, rule[1]) &&
			slices.Index(update, rule[0]) > slices.Index(update, rule[1]) {
			return false
		}
	}

	return true
}

func main() {
	f, err := os.Open("realInput.txt")
	if err != nil {
		panic(err)
	}

	rules := make([][]int, 0)

	total := 0
	scanner := bufio.NewScanner(f)
	isFirstPart := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isFirstPart = false
			continue
		}

		if isFirstPart {
			pages := strings.Split(line, "|")
			page1, _ := strconv.Atoi(pages[0])
			page2, _ := strconv.Atoi(pages[1])
			rules = append(rules, []int{page1, page2})
			continue
		}

		updateLine := strings.Split(line, ",")
		update := make([]int, 0)
		for _, u := range updateLine {
			i, _ := strconv.Atoi(u)
			update = append(update, i)
		}

		if !isUpdateOrdered(rules, update) {
			slices.SortFunc(update, func(a, b int) int {
				for _, rule := range rules {
					if rule[0] == a && rule[1] == b {
						return -1
					}

					if rule[0] == b && rule[1] == a {
						return 1
					}
				}
				return 0
			})
			total += update[len(update)/2]
		}
	}
	fmt.Println(total)
}
