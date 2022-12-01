package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CaloriesPerElf map[int]int

func main() {
	input := readInput()
	inputs := strings.Split(input, "\n")
	elves := caloriesPerElf(inputs)

	biggest := findBiggest(elves)

	fmt.Println("Most calories for an elf:", biggest.sum)

	topThree := sumTopThree(elves)

	fmt.Println("Sum of top three:", topThree)
}

func readInput() string {
	f, _ := os.ReadFile("input.txt")
	i := string(f)
	return i
}

func caloriesPerElf(calorieInput []string) CaloriesPerElf {
	elves := CaloriesPerElf{}

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

func findBiggest(elves CaloriesPerElf) struct{ elf, sum int } {
	biggest := struct{ elf, sum int }{}
	for elf, sum := range elves {
		if sum > biggest.sum {
			biggest.sum = sum
			biggest.elf = elf
		}
	}
	return biggest
}

func sumTopThree(elves CaloriesPerElf) int64 {
	topThree := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		biggest := findBiggest(elves)
		topThree = append(topThree, biggest.sum)
		delete(elves, biggest.elf)
	}

	var sum int64
	for _, v := range topThree {
		sum += int64(v)
	}
	return sum
}
