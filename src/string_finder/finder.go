package string_finder

type Finder interface {
	Find(text, pattern string) int
}
