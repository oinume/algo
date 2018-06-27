package string_finder

import "testing"

func TestBruteForce_Find(t *testing.T) {
	testCases := []struct {
		text       string
		pattern    string
		wantResult int
	}{
		{text: "abcde", pattern: "cd", wantResult: 2},
		{text: "a", pattern: "cd", wantResult: -1},
		{text: "日本語", pattern: "日本", wantResult: 0},
	}

	bf := NewBruteForce()
	for _, tc := range testCases {
		got := bf.Find(tc.text, tc.pattern)
		if tc.wantResult != got {
			t.Errorf("text=%v, pattern=%v: want=%v, got=%v", tc.text, tc.pattern, tc.wantResult, got)
		}
	}
}
