package string_finder

import (
	"unicode/utf8"
)

type BoyerMoore struct {
	shiftTable map[rune]int
}

func NewBoyerMoore() Finder {
	return &BoyerMoore{}
}

func (bm *BoyerMoore) Initialize(pattern string) {
	bm.shiftTable = bm.createShiftTable(pattern)
}

func (bm *BoyerMoore) Find(text string) int {
	return 0
}

func (bm *BoyerMoore) createShiftTable(pattern string) map[rune]int {
	len := utf8.RuneCountInString(pattern)
	table := make(map[rune]int, len)
	for i, r := range pattern {
		// Calculate distance from tail of pattern
		table[r] = len - i - 1
	}
	return table
}
