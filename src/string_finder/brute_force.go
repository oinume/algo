package string_finder

type BruteForce struct {
}

func NewBruteForce() *BruteForce {
	return &BruteForce{}
}

func (bf *BruteForce) Find(text, pattern string) int {
	textRunes := []rune(text)
	patternRunes := []rune(pattern)
	var textPos, patternPos int
	for textPos, patternPos = 0, 0; textPos < len(textRunes) && patternPos < len(patternRunes); {
		// Compare text and pattern char
		if textRunes[textPos] == patternRunes[patternPos] {
			textPos++
			patternPos++
		} else {
			// Not matched
			//
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
