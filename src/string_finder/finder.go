package string_finder

type Finder interface {
	Initialize(pattern string) Finder
	Find(text string) int
}
