package part2

import (
	"bufio"

	"adventofcode/3/utils"
)

func Run(s *bufio.Scanner) (sum int) {
	groups := make([]string, 3)
	iterator := 0
	for s.Scan() {
		row := s.Text()

		groups[iterator] = row

		if iterator == 2 {
			// We have reached third group, let's wrap it up and start over
			commonChar := utils.FindCommonRuneInGroup(groups)
			if commonChar != 0 {
				sum += utils.GetPriorityScore(commonChar)
			}
			iterator = 0
			continue
		}
		iterator++
	}
	return
}
