package main

import (
	"fmt"
	"strings"

	"adventofcode/1/utils"
)

func main() {
	input := utils.ReadInput()
	inputs := strings.Split(input, "\n")

	fmt.Println("Per elf")
	elves := utils.CaloriesPerElf(inputs)

	biggest := make(chan int64)
	go func() {
		b := utils.FindBiggest(elves)
		biggest <- int64(b.Sum)
		close(biggest)
	}()

	topThree := make(chan int64)
	go func() {
		topThree <- utils.SumTopThree(elves)
		close(topThree)
	}()

	fmt.Println("Most calories for an elf:", <-biggest)
	fmt.Println("Sum of top three:", <-topThree)
}
