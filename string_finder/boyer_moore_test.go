package string_finder

import (
	"testing"
)

func TestBoyerMoore_Find(t *testing.T) {
	tests := []struct {
		text       string
		pattern    string
		wantResult int
	}{
		{text: "AAAXABAABBCAC", pattern: "ABBC", wantResult: 7},
		{text: "abcd", pattern: "kcd", wantResult: -1},
		{text: "こんにちは", pattern: "にちは", wantResult: 2},
		{text: "hoge", pattern: "fuga", wantResult: -1},
	}

	for _, test := range tests {
		got := NewBoyerMoore().Initialize(test.pattern).Find(test.text)
		if test.wantResult != got {
			t.Errorf("unexpected result of Find: text=%v, pattern=%v, got=%v", test.text, test.pattern, got)
		}
	}
}
