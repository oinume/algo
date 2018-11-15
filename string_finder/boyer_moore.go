// http://algoful.com/Archive/Algorithm/BMSearch
package string_finder

import (
	"math"
	"unicode/utf8"
)

type BoyerMoore struct {
	pattern    string
	shiftTable map[rune]int
}

func NewBoyerMoore() Finder {
	return &BoyerMoore{}
}

func (bm *BoyerMoore) Initialize(pattern string) Finder {
	bm.pattern = pattern
	bm.shiftTable = bm.createShiftTable(pattern)
	return bm
}

func (bm *BoyerMoore) Find(text string) int {
	textLen := utf8.RuneCountInString(text)
	patternLen := utf8.RuneCountInString(bm.pattern)
	textRunes := []rune(text)
	patternRunes := []rune(bm.pattern)
	// Set start position to the tail of pattern
	textPos := patternLen - 1
	patternPos := 0
	for textPos < textLen {
		//fmt.Printf("text[%v]=%v, pattern[%v]=%v\n",
		//	textPos, string(text[textPos]), patternPos, string(bm.pattern[patternPos]))

		// Set patternPos to the tail of pattern
		patternPos = patternLen - 1
		for patternPos >= 0 && textPos < textLen {
			if textRunes[textPos] == patternRunes[patternPos] {
				// character is matched
				textPos--
				patternPos--
			} else {
				break
			}
		}
		if patternPos < 0 {
			// All characters are matched
			return textPos + 1
		}

		// forward textPos by referring shift table If not match.
		// ただし、今比較した位置より後の位置とする
		p := patternRunes[patternPos]
		var shift1 int
		if v, ok := bm.shiftTable[rune(p)]; ok {
			shift1 = v
		} else {
			shift1 = patternLen
		}
		shift2 := patternLen - patternPos
		textPos += int(math.Max(float64(shift1), float64(shift2)))
	}
	return -1
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
