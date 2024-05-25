package main

import (
	"fmt"
	"slices"
	"strings"
)

func IsPalindrome(str string) bool {
	lc := strings.ToLower(str)

	for i := 0; i < len(lc)/2; i++ {
		if lc[i] != lc[len(lc)-i-1] {
			return false
		}
	}

	return true
}

func IsPalindromeStrings(str string) bool {
	lc := strings.ToLower(str)

	sr := []rune(lc)
	slices.Reverse(sr)

	return string(sr) == lc
}

func IsPalindromeRecursive(str string) bool {
	lc := strings.ToLower(str)

	if len(lc) <= 1 {
		return true
	}

	if lc[0] != lc[len(lc)-1] {
		return false
	}

	return IsPalindromeRecursive(lc[1 : len(lc)-1])
}

func main() {

	in := "Anna"

	fmt.Print(IsPalindromeRecursive(in))
}
