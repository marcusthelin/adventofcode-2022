package utils

import (
	"os"
	"strings"
	"unicode"
)

// Find common character in two rucksacks
func FindCommonRune(first, second string) rune {
	var c rune
	for _, char := range first {
		hasChar := strings.ContainsRune(second, char)
		if hasChar {
			c = char
			break
		}
	}
	return c
}

func FindCommonRuneInGroup(group []string) rune {
	var c rune
	for _, char := range group[0] {
		hasChar := true
		for _, row := range group[1:] {
			hasChar = hasChar && strings.ContainsRune(row, char)
		}
		if hasChar {
			c = char
			break
		}
	}
	return c
}

func ReadFile(path string) *os.File {
	f, _ := os.Open(path)
	return f
}

func GetPriorityScore(char rune) int {
	if unicode.IsLower(char) {
		return int(char) - 96
	}
	return int(char) - 38
}
