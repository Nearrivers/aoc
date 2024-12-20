package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	children map[string]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
		},
	}
}

func (t *Trie) Insert(towel string) {
	node := t.root
	for _, char := range towel {
		str := string(char)
		if node.children[str] == nil {
			node.children[str] = &Node{
				children: make(map[string]*Node),
			}
		}
		node = node.children[str]
	}
	node.isEnd = true
}

func (t *Trie) Search(towel string) bool {
	node := t.root
	for _, char := range towel {
		str := string(char)
		if node.children[str] == nil {
			return false
		}
		node = node.children[str]
	}
	return node.isEnd
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strfile := strings.Split(string(b), "\r\n")
	towels := strings.Split(strfile[0], ", ")

	pattern := `^(`

	for i := range towels {
		if i == len(towels)-1 {
			pattern += towels[i]
			continue
		}
		pattern += towels[i] + "|"
	}
	pattern += `)+$`

	validCount := 0
	reg := regexp.MustCompile(pattern)
	for _, design := range strfile[2:] {
		str := reg.FindAllString(design, -1)
		if len(str) > 0 {
			validCount++
		}
	}

	fmt.Println(validCount)
	return

	trie := NewTrie()
	for _, t := range towels {
		trie.Insert(t)
	}

	// Préfix déjà vus. true si le pattern de la serviette existe, false dans le cas contraire
	seen := make(map[string]bool, 0)

	for _, design := range strfile[2:] {
		start, end := 0, 1
		isPresent := false
		longestValidPrefix := ""
		var lastChunk string

	inner:
		for start < len(design) {
			var chunk string

			if end > len(design) {
				chunk = design[start:]
			} else {
				chunk = design[start:end]
			}

			if chunk == lastChunk {
				break inner
			}

			lastChunk = chunk
			c, ok := seen[chunk]
			// Si pattern déjà vu
			if ok {
				// Et pattern existe dans nos serviettes
				if c {
					isPresent = true
					longestValidPrefix = chunk
					end++
					continue inner
				}

				isPresent = false
				// Si pattern n'existe pas dans nos serviettes
				// et qu'aucun préfix n'a été trouvé
				if len(longestValidPrefix) == 0 {
					end++
					continue inner
				}

				// Si pattern n'existe pas et préfixe précédemment trouvé
				start += len(longestValidPrefix)
				end = start + 1
				longestValidPrefix = ""
				continue inner
			}

			isPresent = trie.Search(chunk)
			if isPresent {
				seen[chunk] = true
				longestValidPrefix = chunk
				end++
				continue inner
			}

			seen[chunk] = false
			if end == len(design) {
				start += len(longestValidPrefix)
				end = start + 1
				longestValidPrefix = ""
				continue
			}

			end++
		}

		if isPresent {
			validCount++
		}
	}

	fmt.Println(validCount)
}
