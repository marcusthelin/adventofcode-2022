package main

import (
	"log"
	"os"
	"sync"

	"adventofcode/2/part1"
	"adventofcode/2/part2"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		f, err := os.Open("input.txt")
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		part1.Run(f)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		f, err := os.Open("input.txt")
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		part2.Run(f)
		wg.Done()
	}()

	wg.Wait()
}
