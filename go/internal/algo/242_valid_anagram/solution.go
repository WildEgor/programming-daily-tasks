package valid_anagram

/**
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
@link https://leetcode.com/problems/valid-anagram
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[rune]int{}
	for _, char := range s {
		sMap[char] += 1
	}

	for _, char := range t {
		value, ok := sMap[char]

		if ok != true || value == 0 {
			return false
		}

		sMap[char]--
	}

	return true
}

func isAnagramAlt(s string, t string) bool {
	c := make([]int, 26) // allocate slice with space for english alphabet
	for _, v := range s {
		i := int(v - 'a') // get char code index part
		c[i]++            // count and decrease later for comparison
	}
	for _, v := range t {
		i := int(v - 'a')
		c[i]--
	}
	for i := range c {
		if c[i] != 0 {
			return false
		}
	}
	return true
}

var Solution = isAnagram
