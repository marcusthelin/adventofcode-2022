package utils

import (
	"os"
	"strconv"
)

// Map containing elf number and number of calories
type ElfCals map[int]int

func ReadInput() string {
	f, _ := os.ReadFile("input.txt")
	i := string(f)
	return i
}

func CaloriesPerElf(calorieInput []string) ElfCals {
	elves := ElfCals{}

	currSum := 0
	currElf := 1
	for _, inp := range calorieInput {
		if inp == "" {
			elves[currElf] = currSum
			currSum = 0
			currElf++
			continue
		}
		val, _ := strconv.ParseInt(inp, 10, 64)
		currSum += int(val)
	}
	return elves
}

func FindBiggest(elves ElfCals) (biggest struct {
	Elf int
	Sum int
},
) {
	for elf, sum := range elves {
		if sum > biggest.Sum {
			biggest.Sum = sum
			biggest.Elf = elf
		}
	}
	return
}

func SumTopThree(e ElfCals) int64 {
	topThree := make([]int, 0, 3)
	elves := make(ElfCals, len(e))

	// Maps are passed by reference, so we need to make a copy
	for elf, sum := range e {
		elves[elf] = sum
	}

	for i := 0; i < 3; i++ {
		biggest := FindBiggest(e)
		topThree = append(topThree, biggest.Sum)
		delete(elves, biggest.Elf)
	}

	var sum int64
	for _, v := range topThree {
		sum += int64(v)
	}
	return sum
}
