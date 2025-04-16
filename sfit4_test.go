package sift4_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ndx-technologies/sift4"
)

func ExampleDistance() {
	d := sift4.Distance("kitten", "sitting", 100, 5, nil)
	fmt.Print(d)
	// Output: 3
}

func ExampleDistance_buffer() {
	var b sift4.Buffer
	d := sift4.Distance("kitten", "sitting", 100, 5, &b)
	fmt.Print(d)
	// Output: 3
}

func TestDistance(t *testing.T) {
	tests := []struct {
		s1          string
		s2          string
		maxOffset   int
		maxDistance int
		distance    int
	}{
		{"kitten", "sitting", 100, 5, 3},
		{"book", "back", 100, 5, 2},
		{"", "abc", 100, 5, 3},
		{"abc", "", 100, 5, 3},
		{"", "", 100, 5, 0},
		{"a", "a", 100, 5, 0},
		{"a", "b", 100, 5, 1},
		{"ab", "abc", 100, 5, 1},
		{"abc", "ab", 100, 5, 1},
		{"abc", "def", 100, 5, 3},
		{"hello", "helo", 100, 5, 1},
		{"world", "word", 100, 5, 1},
		{"halooooxo", "hbloooogo", 100, 5, 6},

		// early exit not reached
		{"distance", "difference", 100, 5, 6},
		{"abcdef", "xyz", 100, 2, 3},
		{"abcdefabcdefabcdefabcdefabcdefabcdef", "xyz", 100, 2, 3},

		// transposition
		{"abc", "acb", 100, 5, 1}, // Damerau–Levenshtein distance, transposition of adjacent characters is one operation
		{"ab", "ba", 100, 5, 1},
		{"abcd", "badc", 100, 5, 2}, // two transpositions
		{"abc", "acb", 100, 5, 1},
		{"aab", "baa", 100, 5, 1},   // covers cursor adjustment when s1[c1] == s2[c2+i]
		{"abcd", "cdab", 100, 5, 2}, // transposition test with cyclic shift
		{"01", "11", 100, 5, 1},
		{"00010", "000010", 100, 5, 2},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			var buf sift4.Buffer
			d := sift4.Distance(tc.s1, tc.s2, tc.maxOffset, tc.maxDistance, &buf)
			if d != tc.distance {
				t.Error(tc, d)
			}
		})
	}
}

func Benchmark_______________________________________(b *testing.B) {}

func BenchmarkSIFT4Distance(b *testing.B) {
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{"empty", "", ""},
		{"one empty", "hello", ""},
		{"equal", "kitten", "kitten"},
		{"different", "kitten", "sitting"},
		{"long different", strings.Repeat("a", 256), strings.Repeat("b", 256)},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for b.Loop() {
				sift4.Distance(tc.s1, tc.s2, 100, 5, nil)
			}
		})
	}

	b.Run("buffer", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				var buffer sift4.Buffer
				for b.Loop() {
					sift4.Distance(tc.s1, tc.s2, 100, 5, &buffer)
				}
			})
		}
	})
}

func FuzzSIFT4Distance(f *testing.F) {
	f.Add("", "")
	f.Add("hello", "")
	f.Add("", "world")
	f.Add("kitten", "sitting")

	f.Fuzz(func(t *testing.T, s1, s2 string) {
		d := sift4.Distance(s1, s2, 100, 5, nil)

		if d < 0 {
			t.Error("d < 0")
		}
		if s1 == s2 && d != 0 {
			t.Error(d)
		}
		if s1 == "" && d != len(s2) {
			t.Error(d)
		}
		if s2 == "" && d != len(s1) {
			t.Error(d)
		}
	})
}
