package string_finder

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
	return nil
}
