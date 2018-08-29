package string_finder

type BoyerMoore struct{}

func NewBoyerMoore() Finder {
	return &BoyerMoore{}
}

func (bm *BoyerMoore) Find(text, pattern string) int {
	return 0
}
