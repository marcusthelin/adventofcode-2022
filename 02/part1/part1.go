package part1

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Map that describes what the different choices can win against
var choices = map[string]map[string]string{
	"opponent": {
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	},
	"me": {
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	},
}

type ChoicePoints struct {
	rock,
	paper,
	scissors int
}
type RoundPoints struct {
	win,
	draw,
	loss int
}

func Run(f io.Reader) {
	s := bufio.NewScanner(f)

	points := map[string]int{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	roundsPoints := RoundPoints{
		win:  6,
		draw: 3,
		loss: 0,
	}
	myPoints := 0
	for s.Scan() {
		// Split the line into two choices
		result := strings.Split(s.Text(), " ")
		var (
			opponentChoice = result[0]
			myChoice       = result[1]
		)

		switch {
		case choices["opponent"][opponentChoice] == choices["me"][myChoice]:
			myPoints += roundsPoints.draw
		case choices["opponent"][opponentChoice] == "rock" && choices["me"][myChoice] == "paper":
			myPoints += roundsPoints.win
		case choices["opponent"][opponentChoice] == "paper" && choices["me"][myChoice] == "scissors":
			myPoints += roundsPoints.win
		case choices["opponent"][opponentChoice] == "scissors" && choices["me"][myChoice] == "rock":
			myPoints += roundsPoints.win
		default:
			myPoints += roundsPoints.loss
		}

		myPoints += points[choices["me"][myChoice]]

	}

	fmt.Println("Part 1:", myPoints)
}
