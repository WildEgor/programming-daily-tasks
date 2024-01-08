package strings_is_subsequence

/**
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
subsequence of a string is a new string that is formed from the original string by deleting some
(can be none) of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
@link https://leetcode.com/problems/is-subsequence/
*/

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	sArr := []rune(s)

	for _, char := range t {
		if len(sArr) == 0 {
			break
		}

		if len(sArr) >= 1 && rune(char) == sArr[0] {
			sArr = sArr[1:]
		}
	}

	return len(sArr) == 0
}

var Solution = isSubsequence
