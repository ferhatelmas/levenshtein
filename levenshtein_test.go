package levenshtein

import "testing"

func log(index int, i interface{}, expected, got int, t *testing.T) {
	t.Errorf("Case %d: %v - expected %d, got %d", index, i, expected, got)
}

func TestMin3(t *testing.T) {
	var tests = []struct {
		a        int
		b        int
		c        int
		expected int
	}{
		{0, 0, 0, 0},
		{1, 2, 3, 1},
		{15129, 123, 123, 123},
		{-1, 12, 2100000000, -1},
	}
	for i, tt := range tests {
		if got := min3(tt.a, tt.b, tt.c); got != tt.expected {
			log(i, []int{tt.a, tt.b, tt.c}, tt.expected, got, t)
		}
	}
}

func TestDist(t *testing.T) {
	var tests = []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"ab", "ab", 0},
		{"ab", "aa", 1},
		{"ab", "aaa", 2},
		{"a", "aa", 1},
		{"bbb", "a", 3},
		{"kitten", "sitting", 3},
		{"aa", "aü", 1},
		{"ferhat", "elmas", 4},
		{"Nico", "nico", 1},
		{"Türkiye", "Atatürk", 7},
	}
	for i, tt := range tests {
		if got := Dist(tt.s1, tt.s2); got != tt.expected {
			log(i, []string{tt.s1, tt.s2}, tt.expected, got, t)
		}
	}
}

func TestDistWithDifferentCosts(t *testing.T) {
	var tests = []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"a", "a", 0},
		{"ab", "ab", 0},
		{"ab", "aa", 3},
		{"ab", "aaa", 5},
		{"aa", "aaa", 2},
		{"a", "bbb", 7},
		{"bbb", "a", 7},
		{"kitten", "sitting", 8},
		{"aa", "aü", 3},
		{"ferhat", "elmas", 11},
		{"Nico", "nico", 3},
		{"Türkiye", "Atatürk", 15},
	}
	c := New(2, 3)
	for i, tt := range tests {
		if got := c.Dist(tt.s1, tt.s2); got != tt.expected {
			log(i, []string{tt.s1, tt.s2}, tt.expected, got, t)
		}
	}
}
