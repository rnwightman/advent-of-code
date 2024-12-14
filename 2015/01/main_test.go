package main

import "testing"

func TestDecodeInstructions(t *testing.T) {
	var tests = []struct {
		input string
		floor int
		pos   int
	}{
		{"(())", 0, 0},
		{"()()", 0, 0},

		{"(((", 3, 0},
		{"(()(()(", 3, 0},

		{"())", -1, 3},

		{")", -1, 1},
		{"()())", -1, 5},
	}

	for _, tc := range tests {
		f, p := DecodeInstructions(tc.input)
		if f != tc.floor || p != tc.pos {
			t.Errorf("DecodeInstructions(%q), expected (%d, %d), got (%d, %d)", tc.input, tc.floor, tc.pos, f, p)
		}
	}
}
