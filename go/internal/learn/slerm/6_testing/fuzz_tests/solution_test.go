package fuzz_tests

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseString(f *testing.F) {
	tcs := []string{"Hello, world", " ", "!12345"}
	for _, tc := range tcs {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := reverseString(orig)
		doubleRev := reverseString(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
