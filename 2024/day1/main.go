package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	colonneDeGauche := make([]int, 0)
	colonneDeDroite := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ligne := strings.Fields(scanner.Text())

		chiffreDeGauche, _ := strconv.Atoi(ligne[0])
		chiffreDeDroite, _ := strconv.Atoi(ligne[1])

		colonneDeGauche = append(colonneDeGauche, chiffreDeGauche)
		colonneDeDroite = append(colonneDeDroite, chiffreDeDroite)
	}

	scoreDeSimilarite := 0
	for i := 0; i < len(colonneDeGauche); i++ {
		chiffreGauche := colonneDeGauche[i]

		cptApparence := 0
		for j := 0; j < len(colonneDeDroite); j++ {
			chiffreDroite := colonneDeDroite[j]

			if chiffreGauche == chiffreDroite {
				cptApparence++
			}
		}

		scoreDeSimilarite += cptApparence * chiffreGauche
	}

	fmt.Println(scoreDeSimilarite)
}
