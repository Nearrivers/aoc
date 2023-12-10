package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var power = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

type hand struct {
	cards string
	bid   int
}

type builtHands []hand

func (h builtHands) Len() int {
	return len(h)
}

func (h builtHands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h builtHands) Less(i, j int) bool {
	iCharCounts := getCharCount(h[i].cards)
	jCharCounts := getCharCount(h[j].cards)

	if iCharCounts.Len() != jCharCounts.Len() {
		return iCharCounts.Len() > jCharCounts.Len()
	}

	var iLessThanJ string = ""

	// On compare les puissances des mains
	for index := 0; index < iCharCounts.Len(); index++ {
		if iCharCounts[index].value < jCharCounts[index].value {
			iLessThanJ = "true"
			break
		}

		if iCharCounts[index].value > jCharCounts[index].value {
			iLessThanJ = "false"
			break
		}
	}

	// Si iLessThanJ est nil cela veut dire que la puissance
	// des deux mains est identique
	// On va donc comparer, carte par carte, la puissance
	if iLessThanJ == "" {
		// On parcourt les cartes
		for index := 0; index < len(h[i].cards); index++ {
			var powerOfi int
			var powerOfj int

			// On tente de convertir la clé du tableau i en int
			key, err := strconv.Atoi(string(h[i].cards[index]))
			if err != nil {
				// Si nous avons une erreur cela veut dire que la clé est une figure
				// donc on se réfère au map power déclaré plus haut
				powerOfi = power[string(h[i].cards[index])]
			} else {
				powerOfi = key
			}

			// Pareil ici mais avec le tableau j
			key, err = strconv.Atoi(string(h[j].cards[index]))
			if err != nil {
				powerOfj = power[string(h[j].cards[index])]
			} else {
				powerOfj = key
			}

			if powerOfi < powerOfj {
				iLessThanJ = "true"
				break
			} else if powerOfi > powerOfj {
				iLessThanJ = "false"
				break
			}
		}
	}

	// Si iLessThanJ est toujours indéfini alors les 2 mains sont identiques
	// en terme de pouvoir ET en terme de carte donc on peut retourner false
	if iLessThanJ == "true" {
		return true
	}

	return false
}

type CharCount struct {
	key   string
	value int
}

type CharCountList []CharCount

func (c CharCountList) Len() int           { return len(c) }
func (c CharCountList) Less(i, j int) bool { return c[i].value > c[j].value }
func (c CharCountList) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func getCharCount(hand string) CharCountList {
	charCount := make(map[string]int)

	// On boucle sur chaque lettre de la main
	for i := 0; i < len(hand); i++ {
		_, ok := charCount[string(hand[i])]

		// Si la lettre y est déjà, on passe l'itération
		// Sinon on initialise à 1
		if ok {
			continue
		} else {
			charCount[string(hand[i])] = 1
		}

		// On refait une boucle sur la même main
		for j := 0; j < len(hand); j++ {
			// Si le char à l'index i est égal au char à l'index j
			// mais que i != j alors on incrémente le compte de 1
			if hand[i] == hand[j] && i != j {
				charCount[string(hand[i])]++
			}
		}
	}

	cl := make(CharCountList, len(charCount))
	i := 0
	for k, v := range charCount {
		cl[i] = CharCount{k, v}
		i++
	}

	// Partie 2 ici
	wildCardCount := 0
	if len(cl) != 1 {
		wildCardIndex := -1
		for i, cc := range cl {
			if cc.key == "J" {
				wildCardCount = cc.value
				wildCardIndex = i
				cc.value = 0
				break
			}
		}

		if wildCardIndex != -1 {
			cl = append(cl[:wildCardIndex], cl[wildCardIndex+1:]...)
		}
	}

	sort.Sort(cl)

	cl[0].value += wildCardCount

	return cl
}

func buildHand(handsWithBids []string) builtHands {
	var buildHands builtHands

	for _, handWithBid := range handsWithBids {
		// handWithBid = strings.Replace(handWithBid, "\t", "", -1)
		h := strings.Split(handWithBid, " ")
		bid, err := strconv.Atoi(h[1])

		if err != nil {
			fmt.Println("erreur de convertion :", err)
		}

		buildHands = append(buildHands, hand{
			cards: h[0],
			bid:   bid,
		})
	}

	return buildHands
}

func main() {
	content, _ := os.ReadFile("input.txt")
	fileContent := string(content)

	handsWithBids := buildHand(strings.Split(fileContent, "\r\n"))
	sort.Sort(handsWithBids)

	totalWinning := 0
	for i, builtHand := range handsWithBids {
		totalWinning += builtHand.bid * (i + 1)
	}

	fmt.Println("Gains totaux :", totalWinning)
}
