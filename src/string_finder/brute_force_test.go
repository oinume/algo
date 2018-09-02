package string_finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBruteForce_Find(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		text       string
		pattern    string
		wantResult int
	}{
		{text: "abcde", pattern: "cd", wantResult: 2},
		{text: "a", pattern: "cd", wantResult: -1},
		{text: "日本語", pattern: "日本", wantResult: 0},
	}

	for _, tc := range testCases {
		bf := NewBruteForce()
		bf.Initialize(tc.pattern)
		got := bf.Find(tc.text)
		a.Equalf(tc.wantResult, got, "text=%v, pattern=%v", tc.text, tc.pattern)
	}
}
