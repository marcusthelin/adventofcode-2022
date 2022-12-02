package part2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var choiceCounters = map[string]string{
	"paper":    "scissors",
	"rock":     "paper",
	"scissors": "rock",
}

var opponentChoiceTranslations = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var outcomePoints = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var choicePoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

func Run(f io.Reader) {
	s := bufio.NewScanner(f)
	myPoints := 0
	for s.Scan() {
		// Split the line into two choices
		result := strings.Split(s.Text(), " ")
		var (
			opponentChoice = opponentChoiceTranslations[result[0]]
			outcome        = result[1]
		)

		// X loss, Y draw, Z win
		switch outcome {
		case "X":
			// Lose
			myChoice := choiceToLose(opponentChoice)
			myPoints += choicePoints[myChoice]
		case "Y":
			// Draw
			myChoice := opponentChoice
			myPoints += choicePoints[myChoice]
		case "Z":
			// Win
			myChoice := choiceCounters[opponentChoice] // Pick the choice that would win against the opponent
			myPoints += choicePoints[myChoice]
		}
		myPoints += outcomePoints[outcome]

	}
	fmt.Println("Part 2:", myPoints)
}

func choiceToLose(opponentChoice string) string {
	if opponentChoice == "rock" {
		return "scissors"
	}
	if opponentChoice == "paper" {
		return "rock"
	}
	if opponentChoice == "scissors" {
		return "paper"
	}
	return ""
}
