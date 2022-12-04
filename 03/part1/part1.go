package part1

import (
	"bufio"

	"adventofcode/3/utils"
)

func Run(s *bufio.Scanner) int {
	sum := 0
	for s.Scan() {
		row := s.Text()
		sum += getPrio(row)
	}
	return sum
}

// Returns the prio score of a row
func getPrio(row string) (sum int) {
	firstRucksack := row[0 : len(row)/2]
	secondRucksack := row[len(row)/2:]

	// fmt.Println(firstRucksack, secondRucksack)
	common := utils.FindCommonRune(firstRucksack, secondRucksack)

	if common == 0 {
		return
	}

	sum += utils.GetPriorityScore(common)

	return
}
