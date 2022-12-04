package main

import (
	"bufio"
	"fmt"
	"sync"

	"adventofcode/3/part1"
	"adventofcode/3/part2"
	"adventofcode/3/utils"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := utils.ReadFile("input.txt")
		defer f.Close()
		s := bufio.NewScanner(f)
		sum := part1.Run(s)
		fmt.Println("Part 1:", sum)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := utils.ReadFile("input.txt")
		defer f.Close()
		s := bufio.NewScanner(f)
		sum := part2.Run(s)
		fmt.Println("Print 2:", sum)
	}()

	wg.Wait()
}
