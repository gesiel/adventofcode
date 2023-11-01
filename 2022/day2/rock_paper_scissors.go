package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// My hand -> X = lose Y = draw Z = win
// Opponent hand -> A = Rock B = Paper C = Scissors
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	handPoints := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	combinationPoints := map[string]int{
		"AX": 3,
		"BX": 1,
		"CX": 2,

		"AY": 1,
		"BY": 2,
		"CY": 3,

		"AZ": 2,
		"BZ": 3,
		"CZ": 1,
	}

	points := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		round := strings.Split(line, " ")

		handScore := handPoints[round[1]]
		roundScore := combinationPoints[strings.Join(round, "")]
		total := handScore + roundScore

		fmt.Printf("Round: %q\nHand score: %v\nRound score: %v\nTotal: %v\n\n", round, handScore, roundScore, total)

		points = points + total
	}

	fmt.Printf("Total: %v\n", points)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
