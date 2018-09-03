package string_finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoyerMoore_Find(t *testing.T) {
	a := assert.New(t)
	testCases := []struct {
		text       string
		pattern    string
		wantResult int
	}{
		{text: "AAAXABAABBCAC", pattern: "ABBC", wantResult: 7},
		{text: "abcd", pattern: "kcd", wantResult: -1},
		{text: "こんにちは", pattern: "にちは", wantResult: 2},
		{text: "hoge", pattern: "fuga", wantResult: -1},
	}

	for _, tc := range testCases {
		got := NewBoyerMoore().Initialize(tc.pattern).Find(tc.text)
		a.Equalf(tc.wantResult, got, "text=%v, pattern=%v", tc.text, tc.pattern)
	}
}
