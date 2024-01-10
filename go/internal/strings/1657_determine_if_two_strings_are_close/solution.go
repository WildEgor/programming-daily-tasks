package determine_if_two_strings_are_close

import (
	"reflect"
	"sort"
)

/**
Two strings are considered close if you can attain one from the other using the following operations:
Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.
Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
@link https://leetcode.com/problems/determine-if-two-strings-are-close/
*/

func counter(word string) (keys, vals [26]int) {
	for i := range word {
		keys[word[i]-'a'] = 1
		vals[word[i]-'a'] += 1
	}
	sort.Ints(vals[:])
	return keys, vals
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	keys1, vals1 := counter(word1)
	keys2, vals2 := counter(word2)

	if reflect.DeepEqual(keys1, keys2) && reflect.DeepEqual(vals1, vals2) {
		return true
	}
	return false
}

var Solution = closeStrings
