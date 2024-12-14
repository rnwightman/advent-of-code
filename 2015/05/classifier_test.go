package main

import (
	"testing"
)

func TestIsNice(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range tests {
		r := IsNice(tc.input)
		if r != tc.want {
			t.Errorf("IsNice(%s): %v, wanted %v", tc.input, r, tc.want)
		}
	}
}

func TestContainsSpacedRepeat(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"xyx", true},
		{"aaa", true},
		{"abcdefeghi", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", true},
	}
	for _, tc := range tests {
		r := containsSpacedRepeat(tc.input)
		if r != tc.want {
			t.Errorf("containsSpacedRepeat(%s): %v, wanted %v", tc.input, r, tc.want)
		}
	}
}

func TestContainsPair(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"xyxy", true},
		{"aaa", false},
		{"aabcdefgaa", true},
	}
	for _, tc := range tests {
		r := containsPair(tc.input)
		if r != tc.want {
			t.Errorf("containsPair(%s): %v, wanted %v", tc.input, r, tc.want)
		}
	}
}
