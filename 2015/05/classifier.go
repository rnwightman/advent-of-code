package main

import (
	"strings"
)

func IsNice(s string) bool {
	if !containsVowels(s) {
		return false
	}
	if !containsDoubleLetter(s) {
		return false
	}
	if containsProhibitted(s) {
		return false
	}

	return true
}

func containsPair(s string) bool {
	for i, j := 0, 1; j < len(s); i, j = i+1, j+1 {
		if s[i] != s[j] {
			continue
		}
	}

	return false
}

func containsSpacedRepeat(s string) bool {
	for i, j := 0, 2; j < len(s); i, j = i+1, j+1 {
		if s[i] == s[j] {
			return true
		}
	}
	return false
}

var vowels = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
}

func containsVowels(s string) bool {
	c := 0
	for _, v := range vowels {
		c += strings.Count(s, v)
	}

	return c >= 3
}

func containsDoubleLetter(s string) bool {
	for i := range s {
		if i == 0 {
			continue
		}

		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func containsProhibitted(s string) bool {
	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
		return true
	}

	return false
}
