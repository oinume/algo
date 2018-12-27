package string_finder

import (
	"testing"
)

func TestBruteForce_Find(t *testing.T) {
	tests := []struct {
		text       string
		pattern    string
		wantResult int
	}{
		{text: "abcde", pattern: "cd", wantResult: 2},
		{text: "a", pattern: "cd", wantResult: -1},
		{text: "日本語", pattern: "日本", wantResult: 0},
	}

	for _, test := range tests {
		got := NewBruteForce().Initialize(test.pattern).Find(test.text)
		if test.wantResult != got {
			t.Errorf("unexpected result of Find: text=%v, pattern=%v, got=%v", test.text, test.pattern, got)
		}
	}
}
