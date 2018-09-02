package string_finder

type Finder interface {
	Initialize(pattern string)
	Find(text string) int
}
