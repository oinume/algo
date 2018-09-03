package string_finder

type BruteForce struct {
	pattern string
}

func NewBruteForce() Finder {
	return &BruteForce{}
}

func (bf *BruteForce) Initialize(pattern string) Finder {
	bf.pattern = pattern
	return bf
}

func (bf *BruteForce) Find(text string) int {
	textRunes := []rune(text)
	patternRunes := []rune(bf.pattern)
	var textPos, patternPos int
	for textPos, patternPos = 0, 0; textPos < len(textRunes) && patternPos < len(patternRunes); {
		// Compare text and pattern char
		if textRunes[textPos] == patternRunes[patternPos] {
			// Match
			textPos++
			patternPos++
		} else {
			// Not match
			textPos = textPos - patternPos + 1 // Make textPos advance with 1 step
			patternPos = 0                     // Reset pattern position
		}
	}

	if patternPos == len(patternRunes) {
		return textPos - patternPos
	} else {
		return -1
	}
}
